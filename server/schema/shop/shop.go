package schema

type Category struct {
	Model
	Name
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	Model
	Name
	Categories []Category `gorm:"many2many:product_categories;"`
}
