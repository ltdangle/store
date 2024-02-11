package models

type Product struct {
	*BaseProduct
	Fields     []*ProductField `gorm:"foreignKey:ProductID"`
	CartItemID uint
}

func NewProduct(baseProduct *BaseProduct) *Product {
	return &Product{BaseProduct: baseProduct}
}

type ProductField struct {
	*BaseProductField
	ProductID uint
}

func NewProductField(b *BaseProductField) *ProductField {
	return &ProductField{BaseProductField: b}
}
