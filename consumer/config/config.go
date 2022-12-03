package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Configurations struct {
	Database DatabaseConfigurations
	Kafka    KafkaConfigurations
}

type DatabaseConfigurations struct {
	Dsn  string
	Pool int
}

type KafkaConfigurations struct {
	Broker string
	Topic  string
}

func LoadConfig(logger *zap.SugaredLogger) *Configurations {
	// err := godotenv.Load()

	// if err != nil {
	// 	logger.Fatalf("Error loading .env file")
	// }

	pool, _ := strconv.Atoi(os.Getenv("POOL"))
	var databaseConfig = &DatabaseConfigurations{
		Dsn:  os.Getenv("DSN"),
		Pool: pool,
	}
	var kafkaConfig = &KafkaConfigurations{
		Broker: os.Getenv("BROKER"),
		Topic:  os.Getenv("TOPIC"),
	}

	var configuration = &Configurations{
		Database: *databaseConfig,
		Kafka:    *kafkaConfig,
	}

	return configuration
}
