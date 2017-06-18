package schema

type User struct {
	Model
	UserName string `json:"user_name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Leads    []Lead `json:"leads" gorm:"many2many:user_leads;"`
}
