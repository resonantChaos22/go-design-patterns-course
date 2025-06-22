package main

import "fmt"

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(src, message string) {
	for _, person := range c.people {
		if person.Name != src {
			person.Receive(src, message)
		}
	}
}
func (c *ChatRoom) Message(src, dst, message string) bool {
	for _, person := range c.people {
		if person.Name == dst {
			person.Receive(src, message)
			return true
		}
	}
	return false
}
func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func (c *ChatRoom) Leave(p *Person) {
	leaveMsg := p.Name + " leaves the chat"

	// Remove person from people slice
	for i, person := range c.people {
		if person == p {
			c.people = append(c.people[:i], c.people[i+1:]...)
			break
		}
	}
	c.Broadcast("Room", leaveMsg)
	p.Room = nil
}

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}
func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}
func (p *Person) PrivateMessage(who, message string) {
	res := p.Room.Message(p.Name, who, message)
	if !res {
		p.Room.Message("Room", "Simon", fmt.Sprintf("No participant named %s", who))
	}
}
func (p *Person) Leave() {
	p.Room.Leave(p)
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func TestChatRoom() {
	room := ChatRoom{}
	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)
	john.Say("Hey there, what's up")
	jane.PrivateMessage("John", "Hello John, how you doin")
	simon := NewPerson("Simon")
	room.Join(simon)
	john.Leave()
	simon.Say("Hey Everyone")
	simon.PrivateMessage("John", "Hey John")
}
