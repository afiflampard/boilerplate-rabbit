package product

import (
	"context"
	"encoding/json"

	"github.com/afiflampard/boilerplate-domain/infra/logger"
	"github.com/afiflampard/boilerplate-domain/infra/rabbit"
)

type ProductMutation interface {
	CreateProductSendEvent(ctx context.Context, product *ProductInput) error
	CreateProductBatch(ctx context.Context, product *ProductInput) error
}

type ProductMutationImpl struct {
	repo     ProductRepository
	rabbitmq *rabbit.RabbitMQ
}

func NewProductMutationImpl(repo ProductRepository, rabbitmq *rabbit.RabbitMQ) ProductMutation {
	return &ProductMutationImpl{repo: repo, rabbitmq: rabbitmq}
}

func (pm *ProductMutationImpl) CreateProductBatch(ctx context.Context, product *ProductInput) error {
	return pm.repo.Create(product)
}

func (pm *ProductMutationImpl) CreateProductSendEvent(ctx context.Context, product *ProductInput) error {
	body, err := json.Marshal(product)
	if err != nil {
		logger.Error("failed to marshal product: %v", err)
		return err
	}

	return pm.rabbitmq.Publish("product", "product.created", body)
}
