package main

import (
	"log"

	"github.com/afiflampard/boilerplate-consumer/cmd/config"
	"github.com/afiflampard/boilerplate-consumer/cmd/consumer"
	"github.com/afiflampard/boilerplate-consumer/cmd/service"
	"github.com/afiflampard/boilerplate-domain/infra/gorm"
	"github.com/afiflampard/boilerplate-domain/infra/rabbit"
	"github.com/afiflampard/boilerplate-domain/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	cfg := config.LoadConfig()
	h := server.New(server.WithHostPorts(":" + cfg.AppPort))

	conn, err := rabbit.Connect(cfg.Rabbit)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	db, err := gorm.InitPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	repository := product.NewProductRepositoryImpl(db)
	mutation := product.NewProductMutationImpl(repository, conn)

	consumer := consumer.NewConsumer(service.NewProductService(mutation))
	go func() {
		if err := consumer.ListenAllQueues(conn.Connection(), cfg); err != nil {
			log.Fatalf("❌ Failed to listen to queues: %v", err)
		}
	}()

	// Jalankan HTTP server Hertz
	h.Spin()
}
