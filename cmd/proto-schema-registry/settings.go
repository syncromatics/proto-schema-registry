package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type settings struct {
	ZookeeperHost string
	Port          int
}

func getSettingsFromEnv() (*settings, error) {
	allErrors := []string{}

	zookeeperHost, ok := os.LookupEnv("ZOOKEEPER_HOST")
	if !ok {
		allErrors = append(allErrors, "ZOOKEEPER_HOST")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		allErrors = append(allErrors, "PORT")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		allErrors = append(allErrors, fmt.Sprintf("failed to convert %s to int", port))
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("Missing required environment variables: %s", strings.Join(allErrors, ", "))
	}

	return &settings{
		ZookeeperHost: zookeeperHost,
		Port:          portInt,
	}, nil
}

func (s *settings) String() string {
	return fmt.Sprintf("ZOOKEEPER_HOST: '%s'\nPORT: %d\n", s.ZookeeperHost, s.Port)
}
