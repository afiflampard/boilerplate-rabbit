package config

import (
	"log"

	"github.com/afiflampard/boilerplate-domain/infra/gorm"
	"github.com/afiflampard/boilerplate-domain/infra/rabbit"
	"github.com/afiflampard/boilerplate-domain/product"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort  string                `envconfig:"APP_PORT" default:"3000"`
	Rabbit   rabbit.RabbitMQConfig `envconfig:"RABBITMQ"`
	Postgres gorm.PostgresConfig   `envconfig:"POSTGRES"`
	Product  product.ProductConfig `envconfig:"PRODUCT"`
}

func LoadConfig() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("❌ Failed to load environment: %v", err)
	}
	return cfg
}
