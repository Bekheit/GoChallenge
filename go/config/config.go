package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"
	"strings"
)

type Configurations struct {
	Server   ServerConfigurations   `koanf:"server"`
	Database DatabaseConfigurations `koanf:"database"`
	Kafka    KafkaConfigurations    `koanf:"kafka"`
}

type ServerConfigurations struct {
	Port int `koanf:"port""`
}

type DatabaseConfigurations struct {
	Dsn  string `koanf:"dsn"`
	Pool int    `koanf:"pool"`
}

type KafkaConfigurations struct {
	Broker string `koanf:"broker"`
	Topic  string `koanf:"topic"`
}

func LoadConfig(logger *zap.SugaredLogger) *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yml"), yaml.Parser())

	if err != nil {
		logger.Fatalf("Failed to locate configurations. %v", err)
	}

	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)

	if err != nil {
		logger.Fatalf("Failed to replace environment variables. %v", err)
	}

	var configuration Configurations

	err = k.Unmarshal("", &configuration)

	if err != nil {
		logger.Fatalf("Failed to load configurations. %v", err)
	}

	return &configuration
}
