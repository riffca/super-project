package models

type Lead struct {
	Model
	Messages []Message
	Members  []User `gorm:"many_to_many:user_leads"`
}

type Message struct {
	Model
	Text   string
	Author User
	Lead   Lead
}
