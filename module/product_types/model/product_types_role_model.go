package model

// ProductTypes Struct
type ProductTypes struct {
	ProductTypesID   int64  `json:"product_types_id"`
	ProductTypesName string `json:"product_types_name"`
}

// ProductTypesList list
type ProductTypesList []ProductTypes

// NewProductTypes func
func NewProductTypes() *ProductTypes {
	return &ProductTypes{}
}
