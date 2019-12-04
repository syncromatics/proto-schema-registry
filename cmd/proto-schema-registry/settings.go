package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type settings struct {
	KafkaBroker       string
	Port              int
	ReplicationFactor int16
}

func getSettingsFromEnv() (*settings, error) {
	allErrors := []string{}

	broker, ok := os.LookupEnv("KAFKA_BROKER")
	if !ok {
		allErrors = append(allErrors, "KAFKA_BROKER")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		allErrors = append(allErrors, "PORT")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		allErrors = append(allErrors, fmt.Sprintf("failed to convert %s to int", port))
	}

	rfInt := 3
	rf, ok := os.LookupEnv("REPLICATION_FACTOR")
	if ok {
		rfInt, err = strconv.Atoi(rf)
		if err != nil {
			allErrors = append(allErrors, fmt.Sprintf("failed to convert %s to int", rf))
		}
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("Missing required environment variables: %s", strings.Join(allErrors, ", "))
	}

	return &settings{
		KafkaBroker:       broker,
		Port:              portInt,
		ReplicationFactor: int16(rfInt),
	}, nil
}

func (s *settings) String() string {
	return fmt.Sprintf("KAFKA_BROKER: '%s'\nPORT: %d\n", s.KafkaBroker, s.Port)
}
