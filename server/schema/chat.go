package schema

type Lead struct {
	Model
	Members    []User `gorm:"many2many:user_leads;"`
	Messages   []Message
	Adress     User
	StatusCode uint
}

type Message struct {
	Model
	Text   string
	LeadId uint
	Sender User
}
