package schema

type User struct {
	Model
	UserName string
	Email    string `gorm:"unique"`
}
