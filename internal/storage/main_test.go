package storage_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"

	client "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/network"
)

var (
	kafkaBroker string
)

func TestMain(m *testing.M) {
	setup()

	result := m.Run()

	teardown()

	os.Exit(result)
}

func setup() {
	var err error
	kafkaBroker, err = setupKafka("storage_test")
	if err != nil {
		panic(err)
	}
}

func teardown() {
	teardownKafka("storage_test")
}

func setupKafka(testName string) (string, error) {
	os.Setenv("DOCKER_API_VERSION", "1.35")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	err = pullImage(cli, "docker.io/confluentinc/cp-schema-registry:5.0.0")
	if err != nil {
		return "", err
	}
	err = pullImage(cli, "docker.io/confluentinc/cp-kafka:5.0.0")
	if err != nil {
		return "", err
	}
	err = pullImage(cli, "docker.io/confluentinc/cp-zookeeper:5.0.0")
	if err != nil {
		return "", err
	}

	removeContainers(cli, testName)

	zookeeperIP, err := createZookeeperContainer(cli, testName)
	if err != nil {
		return "", err
	}

	brokerIP, err := createKafkaContainer(cli, zookeeperIP, testName)
	if err != nil {
		return "", err
	}

	err = waitForKafkaToBecomeAvailable(brokerIP)
	if err != nil {
		return "", errors.Wrap(err, "failed to wait for broker online")
	}

	return brokerIP, nil
}

func teardownKafka(testName string) {
	os.Setenv("DOCKER_API_VERSION", "1.35")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	removeContainers(cli, testName)
}

func pullImage(client *client.Client, image string) error {
	r, err := client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return nil
}

func removeContainers(client *client.Client, testName string) {
	client.ContainerRemove(context.Background(), fmt.Sprintf("%s_zookeeper_test", testName), types.ContainerRemoveOptions{Force: true})
	client.ContainerRemove(context.Background(), fmt.Sprintf("%s_kafka_test", testName), types.ContainerRemoveOptions{Force: true})
	client.ContainerRemove(context.Background(), fmt.Sprintf("%s_registry_test", testName), types.ContainerRemoveOptions{Force: true})
}

func createZookeeperContainer(cli *client.Client, testName string) (string, error) {
	config := container.Config{
		Image: "docker.io/confluentinc/cp-zookeeper:5.0.0",
		Env: []string{
			"ZOOKEEPER_CLIENT_PORT=2181",
		}}

	hostConfig := container.HostConfig{}

	networkConfig := network.NetworkingConfig{}

	create, err := cli.ContainerCreate(context.Background(), &config, &hostConfig, &networkConfig, fmt.Sprintf("%s_zookeeper_test", testName))
	if err != nil {
		return "", err
	}

	conChan, errChan := cli.ContainerWait(context.Background(), create.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errChan:
		return "", err
	case <-conChan:
	}

	err = cli.ContainerStart(context.Background(), create.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	inspect, err := cli.ContainerInspect(context.Background(), create.ID)
	if err != nil {
		return "", err
	}

	return inspect.NetworkSettings.IPAddress, nil
}

func createKafkaContainer(cli *client.Client, zookeeperIP string, testName string) (string, error) {
	config := container.Config{
		Image: "docker.io/confluentinc/cp-kafka:5.0.0",
		Env: []string{
			"HOST_IP=kafka",
			"KAFKA_NUM_PARTITIONS=10",
			"KAFKA_DEFAULT_REPLICATION_FACTOR=1",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
			"KAFKA_REPLICATION_FACTOR=1",
			"KAFKA_BROKER_ID=1",
			fmt.Sprintf("KAFKA_ZOOKEEPER_CONNECT=%s", zookeeperIP),
		},
		Entrypoint: nil,
		Cmd: []string{
			"/bin/sh",
			"-c",
			"export KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://$(hostname -i):9092 && exec /etc/confluent/docker/run",
		},
	}

	hostConfig := container.HostConfig{}

	networkConfig := network.NetworkingConfig{}

	create, err := cli.ContainerCreate(context.Background(), &config, &hostConfig, &networkConfig, fmt.Sprintf("%s_kafka_test", testName))
	if err != nil {
		return "", err
	}

	conChan, errChan := cli.ContainerWait(context.Background(), create.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errChan:
		return "", err
	case <-conChan:
	}

	err = cli.ContainerStart(context.Background(), create.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	inspect, err := cli.ContainerInspect(context.Background(), create.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:9092", inspect.NetworkSettings.IPAddress), nil
}

func waitForKafkaToBecomeAvailable(broker string) error {
	deadline := time.Now().Add(120 * time.Second)
	var lastError error
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Admin.Timeout = 10 * time.Second
	config.Net.DialTimeout = 10 * time.Second

	for time.Now().Before(deadline) {
		admin, err := sarama.NewClusterAdmin([]string{broker}, config)
		if err != nil {
			lastError = err
			time.Sleep(100 * time.Millisecond)
			continue
		}
		defer admin.Close()

		_, err = admin.ListTopics()
		if err != nil {
			lastError = err
			continue
		}

		return nil
	}

	return lastError
}
