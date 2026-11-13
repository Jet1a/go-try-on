package db

import (
	"github.com/Jet1a/go-try-on/api/internal/product"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error {

	err := db.Exec(`
	  DO $$
        BEGIN
            CREATE TYPE product_category AS ENUM (
                'shirt',
                'pants',
                'shoes'
            );
        EXCEPTION
            WHEN duplicate_object THEN NULL;
        END
        $$;
	`).Error

	if err != nil {
		return err
	}

	return db.AutoMigrate(
		&product.Product{},
		// &product.ProductCategory{},
		&product.ProductSize{},
	)
}
