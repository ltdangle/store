package models

type Product struct {
	*BaseProduct
	CartItemUuid string
}

func NewProduct() *Product {
	return &Product{BaseProduct: NewBaseProduct()}
}

type ProductField struct {
	*BaseProductField
	ProductUuid string 
}

func NewProductField() *ProductField {
	return &ProductField{BaseProductField: NewBaseProductField()}
}
