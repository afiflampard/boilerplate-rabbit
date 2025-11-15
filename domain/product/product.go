package product

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func CreateProduct(input ProductInput) *Product {
	id := uuid.New()
	return &Product{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
}
