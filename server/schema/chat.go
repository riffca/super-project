package schema

type Lead struct {
	Model
	Members    []User    `json:"members" gorm:"many2many:user_leads;"`
	Messages   []Message `json:"messages"`
	CreatedBy  User      `json:"created_by" gorm:"ForeignKey:ID;AssociationForeignKey:CreatorID;"`
	CreatorID  uint64    `json:"creator_id"`
	Adress     User      `json:"adress" gorm:"ForeignKey:ID;AssociationForeignKey:AdressID;"`
	AdressID   uint64    `json:"adress_id"`
	StatusCode uint64    `json:"status_code"`
}

type Member struct {
	User User
	Role uint64
}

type Message struct {
	Model
	Text     string
	LeadId   uint64
	Sender   User
	SenderID uint64
}
