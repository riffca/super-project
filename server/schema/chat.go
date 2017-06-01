package schema

type Lead struct {
	Model
	UserID uint
}

type Conversation struct {
	Model
	Memebers  []User
	Messagges []Message
}

type Message struct {
	Model
	Text           string
	ConversationID uint
}
