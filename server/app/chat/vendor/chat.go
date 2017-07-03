package vendor

import "fmt"

type Chat struct {
	Conversations map[string][]string
}

func (c *Chat) CreateConversation(sessionID string) {
	i := []string{sessionID}
	c.Conversations[sessionID] = i
	fmt.Println("ALL CONVERSATIONS", c.Conversations)
}

func (c *Chat) RemoveConversation(sessionID string) {
	delete(c.Conversations, sessionID)
}

func (c *Chat) CheckAdress(adress string, member string) bool {

	var check bool = false
	for k := range c.Conversations {
		if adress == k {
			check = true
			// for _, v := range v {
			//  if v == member {
			//    return true
			//  }
			// }
		}
	}

	fmt.Println(check)
	return check

}

func newChat() *Chat {
	return &Chat{Conversations: map[string][]string{"default": {"default"}}}
}
