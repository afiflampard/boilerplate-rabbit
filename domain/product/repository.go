package product

import "gorm.io/gorm"

type ProductRepository interface {
	Create(product *ProductInput) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (pr *ProductRepositoryImpl) Create(product *ProductInput) error {
	return pr.db.Create(product).Error
}
