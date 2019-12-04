package storage

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/syncromatics/proto-schema-registry/pkg/protobuf"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// FileStorage is a registry storage driver that keeps schemas in file storage
type FileStorage struct {
	consumer      *kafka
	rootDirectory string

	mtx               sync.Mutex
	maxID             uint32
	idMap             map[uint32]string
	topicHash         map[string]map[string]uint32
	latestTopicSchema map[string]uint32
}

// NewFileStorage creates a new file storage driver
func NewFileStorage(broker string, replicas int16, rootDirectory string, topic string, timeout time.Duration) (*FileStorage, error) {
	file := &FileStorage{
		rootDirectory:     rootDirectory,
		topicHash:         map[string]map[string]uint32{},
		idMap:             map[uint32]string{},
		latestTopicSchema: map[string]uint32{},
	}

	consumer, err := newKafkaConsumer(broker, replicas, file, topic, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kafka consumer")
	}

	file.consumer = consumer

	return file, nil
}

// Run runs the file storage driver and will fail if the connection to kafka fails
func (f *FileStorage) Run(ctx context.Context) func() error {
	return func() error {
		c, cancel := context.WithCancel(ctx)
		grp, c := errgroup.WithContext(c)

		grp.Go(f.consumer.Run(c))

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

// GetSchema will attempt to get a schema by id
func (f *FileStorage) GetSchema(ctx context.Context, id uint32) ([]byte, bool, error) {
	err := f.waitForConsumerCatchup(30 * time.Second)
	if err != nil {
		return nil, false, err
	}

	f.mtx.Lock()
	defer f.mtx.Unlock()

	file, err := f.getSchema(id)
	if err != nil {
		return nil, false, err
	}

	return file, true, nil
}

// RegisterSchema will register a schema or return the id if it already exists
func (f *FileStorage) RegisterSchema(ctx context.Context, topic string, schema []byte) (uint32, []string, bool, error) {
	err := f.waitForConsumerCatchup(30 * time.Second)
	if err != nil {
		return 0, nil, false, err
	}

	f.mtx.Lock()
	defer f.mtx.Unlock()

	if t, ok := f.topicHash[topic]; ok {
		md5 := md5.Sum(schema)

		if id, ok := t[fmt.Sprintf("%x", md5)]; ok {
			return id, nil, true, nil
		}
	}

	if current, ok := f.latestTopicSchema[topic]; ok {
		file, err := f.getSchema(current)
		if err != nil {
			return 0, nil, false, errors.Wrap(err, "failed to get current schema")
		}

		compErrors, ok, err := f.compareSchemas(file, schema)
		if err != nil {
			return 0, nil, false, errors.Wrap(err, "failed comparing schemas")
		}

		if !ok {
			return 0, compErrors, false, nil
		}
	}

	err = protobuf.Validate(schema)
	if err != nil {
		return 0, []string{err.Error()}, false, nil
	}

	f.maxID++

	err = f.consumer.Produce(ctx, &Schema{
		Subject: topic,
		Schema:  schema,
		ID:      f.maxID,
	})
	if err != nil {
		return 0, nil, false, errors.Wrap(err, "failed to send schema to kafka")
	}

	return f.maxID, nil, true, nil
}

func (f *FileStorage) store(schema *Schema) error {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	d := path.Join(f.rootDirectory, schema.Subject, strconv.Itoa(int(schema.ID)))
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			return errors.Wrap(err, "failed to create schema directory")
		}
	}
	p := path.Join(d, "data")
	err := ioutil.WriteFile(p, schema.Schema, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to save schema to file")
	}

	f.idMap[schema.ID] = p
	f.maxID = schema.ID

	if _, ok := f.topicHash[schema.Subject]; !ok {
		f.topicHash[schema.Subject] = map[string]uint32{}
	}

	md5 := md5.Sum(schema.Schema)
	f.topicHash[schema.Subject][fmt.Sprintf("%x", md5)] = schema.ID

	f.latestTopicSchema[schema.Subject] = schema.ID

	return nil
}

func (f *FileStorage) getSchema(id uint32) ([]byte, error) {
	s, ok := f.idMap[id]
	if !ok {
		return nil, errors.Errorf("failed to find schema for '%d'", id)
	}

	file, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read schema file")
	}

	return file, nil
}

func (f *FileStorage) waitForConsumerCatchup(duration time.Duration) error {
	deadline := time.Now().Add(duration)
	var offset int64
	var err error
	for time.Now().Before(deadline) {
		offset, err = f.consumer.GetLag()
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}

		if offset == 0 {
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}

	if err != nil {
		return errors.Wrap(err, "failed to get lag")
	}

	return errors.Errorf("failed to reach 0 lag. last reported lag '%d'", offset)
}

func (f *FileStorage) compareSchemas(current []byte, next []byte) ([]string, bool, error) {
	cu, err := gzip.NewReader(bytes.NewBuffer(current))
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to create gzip reader for current schema")
	}

	cb, err := ioutil.ReadAll(cu)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to unzip current schema")
	}

	nu, err := gzip.NewReader(bytes.NewBuffer(next))
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to create gzip reader for next schema")
	}

	nb, err := ioutil.ReadAll(nu)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to unzip next schema")
	}

	ok, comErrors, err := protobuf.CheckForBreakingChanges(cb, nb)
	if err != nil {
		return nil, false, errors.Wrap(err, "failed to compare schemas")
	}

	return comErrors, ok, nil
}
