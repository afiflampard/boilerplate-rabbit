package consumer

import (
	"context"
	"encoding/json"

	"github.com/afiflampard/boilerplate-domain/product"
)

func (c *Consumer) CreateProductConsumer(msg []byte) error {
	var product product.ProductInput
	if err := json.Unmarshal(msg, &product); err != nil {
		return err
	}
	return c.service.CreateProduct(context.Background(), &product)
}
