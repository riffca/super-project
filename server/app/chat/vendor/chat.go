package vendor

import "fmt"

type Member struct {
	UserName string
}

type Room struct {
	Name    string
	Members map[string]string
}

type Chat struct {
	Conversations map[string]Room
}

func (c *Chat) CreateConversation(sessionID string) {
	i := Room{"Комната", map[string]string{sessionID: sessionID}}
	c.Conversations[sessionID] = i
	fmt.Println("ALL CONVERSATIONS", c.Conversations)
}

func (c *Chat) RemoveConversation(sessionID string) {
	// delete user from all conversations
	for _, room := range c.Conversations {
		delete(room.Members, sessionID)
	}
	delete(c.Conversations, sessionID)
}

func (c *Chat) CheckAdress(adress string, member string) bool {

	var check bool = false
	for k := range c.Conversations {
		if adress == k {
			check = true
		}
	}

	fmt.Println(check)
	return check

}

func newChat() *Chat {
	return &Chat{
		Conversations: map[string]Room{
			"default": {
				"default", map[string]string{
					"default": "default",
				},
			},
		},
	}
}
