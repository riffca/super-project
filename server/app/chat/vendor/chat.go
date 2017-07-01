package vendor

import "fmt"

type Chat struct {
	conversations map[string][]string
}

func (c *Chat) createConversation(sessionID string) {
	i := []string{sessionID}
	c.conversations[sessionID] = i
	fmt.Println("ALL CONVERSATIONS", c.conversations)
}

func (c *Chat) checkAdress(adress string, member string) bool {

	var check bool = false
	for k := range c.conversations {
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
	return &Chat{conversations: map[string][]string{"default": {"default"}}}
}
