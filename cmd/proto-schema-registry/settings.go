package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type settings struct {
	KafkaBroker           string
	Port                  int
	ReplicationFactor     int16
	SecondsToWaitForKafka int
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

	secondsToWait := 30
	s, ok := os.LookupEnv("SECONDS_TO_WAIT_FOR_KAFKA")
	if ok {
		secondsToWait, err = strconv.Atoi(s)
		if err != nil {
			allErrors = append(allErrors, fmt.Sprintf("failed to convert %s to int", s))
		}
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("Missing required environment variables: %s", strings.Join(allErrors, ", "))
	}

	return &settings{
		KafkaBroker:           broker,
		Port:                  portInt,
		ReplicationFactor:     int16(rfInt),
		SecondsToWaitForKafka: secondsToWait,
	}, nil
}

func (s *settings) String() string {
	return fmt.Sprintf("KAFKA_BROKER: '%s'\nPORT: %d\n", s.KafkaBroker, s.Port)
}
