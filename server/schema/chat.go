package schema

type Lead struct {
	Model
	Members    []User `gorm:"many2many:user_leads;"`
	Messages   []Message
	CreatedBy  User `gorm:"unique"`
	Adress     User `gorm:"unique"`
	StatusCode uint64
}

type Message struct {
	Model
	Text   string
	LeadId uint64
	Sender User
}
