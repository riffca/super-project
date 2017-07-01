/*
PubSub package provides simple mechanism to implement publisher subscriber relation.
 type Timer struct {
  pubsub.Publisher
 }
 timer := new(Timer)
 go func() {
    for {
    time.Sleep(time.Second)
    timer.Publish(time.Now())
  }
 }()
 reader, _ := timer.SubReader()
 for {
  fmt.Println(reader.Read())
 }
Memory considerations: memory consumption increases if subscribers do not consume messages as fast as published by Publisher.
There's no need to unsubscribe explicitelly, once SubReader reference is lost GC takes care of it. The same applies to subscription channel.
You might want to hide Publish method in composition scenarios:
 type Timer struct {
  p pubsub.Publisher
 }
In that case you need to provide access to SubReader and SubChannel methods:
 func (t *Timer) SubReader() (pubsub.Reader, interface{}) {
  return t.p.SubReader()
 }
*/
package vendor

import (
	"fmt"
	"sync"
)

// Subscription Reader is used to read messages published by Publisher
type SubReader interface {
	// Read operation blocks and waits for message from Publisher
	Read() interface{}
	getSession() string
}

// Publisher is used to publish messages. Can be directly created.
type Publisher struct {
	m         sync.Mutex
	lastMsg   *msg
	sessionID string
}

type subscriber struct {
	in      chan *msg
	session string
}

type msgSchema map[string]interface{}

type msg struct {
	val  interface{}
	next chan *msg
}

var chat *Chat = newChat()

func newMsg(val interface{}) *msg { return &msg{val: val, next: make(chan *msg, 1)} }

// Publish publishes a message to all existing subscribers
func (p *Publisher) Publish(val interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	msg := newMsg(val)
	if p.lastMsg != nil {
		p.lastMsg.next <- msg
	}
	p.lastMsg = msg
}

// SubReader returns a new reader for reading published messages and a last published message.
func (p *Publisher) SubReader() (reader SubReader, lastMsg interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.lastMsg == nil {
		p.lastMsg = newMsg(nil)
	}

	chat.createConversation(p.sessionID)

	return &subscriber{p.lastMsg.next, p.sessionID}, p.lastMsg.val
}

// SubChannel returns a new channel for reading published messages and a last published message.
// If published messages equals (==) finalMsg then channel is closed afer putting message into channel.
func (p *Publisher) SubChannel(finalMsg interface{}, session string) (
	msgChan <-chan interface{},
	lastMsg interface{},
) {

	p.sessionID = session
	listener, cur := p.SubReader()
	outch := make(chan interface{})
	go listen(listener, outch, finalMsg)
	return outch, cur

}

func listen(subscriber SubReader, ch chan interface{}, finalMsg interface{}) {
	defer close(ch)
	for {

		state := subscriber.Read()

		if chat.checkAdress(state["adress_id"], state["session_id"]) {
			continue
		}
		// if state != subscriber.getSession() {
		//  continue
		// }

		ch <- state
		if state == finalMsg {
			return
		}
	}
}

func (s *subscriber) Read() interface{} {
	msg := <-s.in
	//fmt.Println("msg := <-s.in:", msg)
	s.in <- msg
	//fmt.Println("s.in <- msg:", s.in)
	s.in = msg.next
	//fmt.Println("s.in = msg.next", s.in)
	return msg.val
}

func (s *subscriber) getSession() string {
	return s.session
}
