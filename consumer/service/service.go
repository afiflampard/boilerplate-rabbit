package service

import (
	"context"

	"github.com/afiflampard/boilerplate-domain/product"
)

type ProductService struct {
	productMutation product.ProductMutation
}

func NewProductService(productMutation product.ProductMutation) ProductService {
	return ProductService{productMutation: productMutation}
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *product.ProductInput) error {
	return ps.productMutation.CreateProductBatch(ctx, product)
}
