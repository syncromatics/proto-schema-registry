package storage

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// SchemaStorer stores the schema locally
type SchemaStorer interface {
	store(schema *Schema) error
}

type kafka struct {
	client         sarama.Client
	replicas       int16
	storer         SchemaStorer
	topic          string
	producer       sarama.SyncProducer
	broker         string
	producedOffset int64

	mtx           sync.RWMutex
	currentOffset int64
}

func newKafkaConsumer(broker string, replicas int16, storer SchemaStorer, topic string, timeout time.Duration) (*kafka, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 1
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.MaxVersion

	var client sarama.Client
	var err error

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		client, err = sarama.NewClient([]string{broker}, config)
		if err == nil {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "timeout waiting to create kafka client")
	}

	p, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kafka producer")
	}

	admin, err := sarama.NewClusterAdmin([]string{broker}, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cluster admin")
	}
	defer admin.Close()

	compact := "compact"
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: replicas,
		ConfigEntries: map[string]*string{
			"cleanup.policy": &compact,
		},
	}
	err = admin.CreateTopic(topic, topicDetail, false)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, errors.Wrap(err, "failed to create schema topic")
	}

	return &kafka{
		client:        client,
		replicas:      replicas,
		storer:        storer,
		topic:         topic,
		producer:      p,
		currentOffset: -1,
		broker:        broker,
	}, nil
}

func (k *kafka) Run(ctx context.Context) func() error {
	return func() error {

		c, cancel := context.WithCancel(ctx)
		grp, c := errgroup.WithContext(c)

		grp.Go(k.consume(c))

		select {
		case <-ctx.Done():
			cancel()
			return nil
		case <-c.Done():
			cancel()
			return grp.Wait()
		}
	}
}

func (k *kafka) GetLag() (int64, error) {
	err := k.client.RefreshMetadata(k.topic)
	if err != nil {
		return -1, errors.Wrap(err, "failed to refresh metadata")
	}

	offset, err := k.client.GetOffset(k.topic, 0, sarama.OffsetNewest)
	if err != nil {
		return -1, errors.Wrap(err, "failed getting latest offset")
	}

	k.mtx.RLock()
	defer k.mtx.RUnlock()
	if k.producedOffset > offset {
		return k.producedOffset - k.currentOffset, nil
	}

	return offset - k.currentOffset - 1, nil
}

func (k *kafka) Produce(ctx context.Context, schema *Schema) error {
	key := keySchema{
		Subject: schema.Subject,
		ID:      schema.ID,
		Magic:   2,
		KeyType: "SCHEMA",
	}

	kb, err := json.Marshal(key)
	if err != nil {
		return errors.Wrap(err, "failed to marshal key")
	}

	vb, err := json.Marshal(schema)
	if err != nil {
		return errors.Wrap(err, "failed to marshal schema")
	}

	_, offset, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.ByteEncoder(kb),
		Value: sarama.ByteEncoder(vb),
	})
	if err != nil {
		return errors.Wrap(err, "failed to send schema to kafka")
	}

	k.producedOffset = offset + 1

	return nil
}

func (k *kafka) consume(ctx context.Context) func() error {
	return func() error {
		config := sarama.NewConfig()
		config.Version = sarama.MaxVersion
		config.Consumer.Return.Errors = true
		config.Consumer.Offsets.Initial = sarama.OffsetOldest

		consumer, err := sarama.NewConsumer([]string{k.broker}, config)
		if err != nil {
			return errors.Wrap(err, "failed to create consumer")
		}
		defer consumer.Close()

		c, err := consumer.ConsumePartition(k.topic, 0, sarama.OffsetOldest)
		if err != nil {
			return errors.Wrap(err, "failed to consume partition")
		}
		defer c.Close()

		for {
			select {
			case <-ctx.Done():
				return nil

			case err := <-c.Errors():
				return errors.Wrap(err, "failed to consume messages")

			case m := <-c.Messages():
				err := k.processMessage(m)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (k *kafka) processMessage(message *sarama.ConsumerMessage) error {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	schema, err := newSchema(message)
	if err != nil {
		return err
	}

	err = k.storer.store(schema)
	if err != nil {
		return errors.Wrap(err, "failed to store schema")
	}

	k.currentOffset = message.Offset

	return nil
}

// Schema is a proto schema
type Schema struct {
	Subject string
	ID      uint32
	Schema  []byte
	Deleted bool
}

type keySchema struct {
	Subject string
	ID      uint32
	Magic   int
	KeyType string
}

func newSchema(message *sarama.ConsumerMessage) (*Schema, error) {
	key := &keySchema{}
	err := json.Unmarshal(message.Key, key)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal '%s' to schema key", string(message.Key))
	}

	value := &Schema{}
	err = json.Unmarshal(message.Value, value)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to demarshal '%s' to schema value", string(message.Value))
	}

	return value, nil
}
