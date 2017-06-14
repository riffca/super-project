package schema

type User struct {
	Model
	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Leads    []Lead `gorm:"many2many:user_leads;"`
}
