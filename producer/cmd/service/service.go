package service

import (
	"context"

	"github.com/afiflampard/boilerplate-domain/product"
)

type ProductService struct {
	productMutaion product.ProductMutation
}

func NewProductService(productMutaion product.ProductMutation) ProductService {
	return ProductService{productMutaion: productMutaion}
}

func (ps *ProductService) CreateProductSendEvent(ctx context.Context, product *product.ProductInput) error {
	return ps.productMutaion.CreateProductSendEvent(ctx, product)
}
