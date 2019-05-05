package model

import (
	"time"

	m "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product_types/model"
)

// Product Struct
type Product struct {
	ProductID    int64          `json:"product_id"`
	ProductTypes m.ProductTypes `json:"product_types"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	CapitalPrice float64        `json:"capital_price"`
	SellingPrice float64        `json:"selling_price"`
	Image        string         `json:"image"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// ProductList list
type ProductList []Product

// NewProduct func
func NewProduct() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
