package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Kafka    KafkaConfigurations
}

type ServerConfigurations struct {
	Port int
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

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	var serverConf = &ServerConfigurations{
		Port: port,
	}

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
		Server:   *serverConf,
		Database: *databaseConfig,
		Kafka:    *kafkaConfig,
	}

	return configuration
}
