package schema

type Page struct {
	Model
	Name    string `gorm:"unique"`
	Content string
}

func (u *Page) AfterSave() (err error) {
	u.Content = `&quot;hello&quot;:&quot;json&quot;`
	return
}
