package main

import (
	"context"
	"log"

	"github.com/afiflampard/boilerplate-domain/infra/gorm"
	"github.com/afiflampard/boilerplate-domain/infra/rabbit"
	"github.com/afiflampard/boilerplate/cmd/config"
	"github.com/afiflampard/boilerplate/cmd/routes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️  No .env file found (using system env)")
	}
	cfg := config.LoadConfig()

	db, err := gorm.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	rabbitMQ, err := rabbit.Connect(cfg.Rabbit)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	h := server.New(server.WithHostPorts(":" + cfg.AppPort))
	routes.SetupRouter(ctx, h, db, rabbitMQ)
	h.Spin()
}
