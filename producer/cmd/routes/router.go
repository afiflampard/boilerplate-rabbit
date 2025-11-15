package routes

import (
	"context"

	"github.com/afiflampard/boilerplate-domain/infra/rabbit"
	"github.com/afiflampard/boilerplate-domain/product"
	"github.com/afiflampard/boilerplate/cmd/handler"
	"github.com/afiflampard/boilerplate/cmd/service"
	"github.com/cloudwego/hertz/pkg/app/server"
	"gorm.io/gorm"
)

func SetupRouter(ctx context.Context, h *server.Hertz, db *gorm.DB, rabbitMQ *rabbit.RabbitMQ) {
	repository := product.NewProductRepositoryImpl(db)
	mutation := product.NewProductMutationImpl(repository, rabbitMQ)
	service := service.NewProductService(mutation)
	handler := handler.NewHandler(service)

	h.POST("/product", handler.CreateProductSendEvent)
}
