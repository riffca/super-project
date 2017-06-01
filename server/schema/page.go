package schema

type Page struct {
	Model
	Name    string `gorm:"unique"`
	Content string
}
