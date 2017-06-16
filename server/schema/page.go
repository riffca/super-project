package schema

type Page struct {
	Model
	Name    string `json:"name" gorm:"unique"`
	Content string `json:content"" `
}

func (u *Page) BeforeCreate() (err error) {
	u.Content = `{&quot;hello&quot;:&quot;json&quot;}`
	return
}
