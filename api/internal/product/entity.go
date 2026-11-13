package product

import (
	"time"
	"uuid"
)

type Product struct {
	ID          uuid.UUID
	Slug        uint
	Name        string
	Description *string // nullable
	Price       float64

	Category  ProductCategory
	Sizes     []ProductSize
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

func (Product) TableName() string {
	return "products"
}
