package schema

type Lead struct {
	Model
	Members    []User    `json:"members" gorm:"many2many:user_leads;"`
	Messages   []Message `json:"messages"`
	CreatedBy  User      `json:"created_by" gorm:"unique"`
	Adress     User      `json:"adress" gorm:"unique"`
	StatusCode uint64    `json:"status_code"`
}

type Member struct {
	ConnectedAt User
	Role        uint64
}

type Message struct {
	Model
	Text   string
	LeadId uint64
	Sender User
}
