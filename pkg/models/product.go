package models

type Product struct {
	*BaseProduct
	Fields     []*ProductField `gorm:"foreignKey:ProductID"`
	CartItemID uint
}

func NewProduct() *Product {
	return &Product{BaseProduct: NewBaseProduct()}
}

type ProductField struct {
	*BaseProductField
	ProductID uint
}

func NewProductField(b *BaseProductField) *ProductField {
	return &ProductField{BaseProductField: b}
}
