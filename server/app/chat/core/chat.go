package core

import "fmt"
import "time"

type Message struct {
	Date   time.Time
	Text   string
	Sender string
}

type Room struct {
	Name     string
	Members  map[string]string
	Messages []Message
}

type ChatCore struct {
	Conversations map[string]Room
}

func (c *ChatCore) CreateConversation(sessionID string) {
	i := Room{
		Name: "Комната",
		Members: map[string]string{
			sessionID: sessionID,
		}}
	c.Conversations[sessionID] = i
	fmt.Println("ALL CONVERSATIONS", len(c.Conversations))
}

func (c *ChatCore) RemoveConversation(sessionID string) {
	// delete user from all conversations
	for _, room := range c.Conversations {
		delete(room.Members, sessionID)
	}
	delete(c.Conversations, sessionID)
}

func (c *ChatCore) CheckAdress(adress string, member string) bool {
	var check bool = false
	for k := range c.Conversations {
		if adress == k {
			check = true
		}
	}
	return check
}

func (c *ChatCore) saveMessage(adress string, sender string, text string) {
	m := c.Conversations[adress]
	m.Messages = append(m.Messages, Message{Text: text, Sender: sender})
	c.Conversations[adress] = m
}

func newChat() *ChatCore {
	return &ChatCore{
		Conversations: map[string]Room{
			"default": {
				Name: "default",
				Members: map[string]string{
					"default": "default",
				},
			},
		},
	}
}
