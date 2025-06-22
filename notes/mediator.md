- `Mediator` is a component that facilitates communication between other components without them necessarily being aware of each other or having direct (reference) access to each other.
- This comes in handy if objects are created and destroyed on regular basis and this could cause a problem if we are relying on pointer to that objects for communication as the pointers can be dead.

## Chat Room

- In a chat room, if we send messages to another person based on the pointer that refers to that person, it could cause an issue if they leave the chat room as the object is destroyed.
- So we can use a Mediator as a `ChatRoom` which is responsible for validating and sending the messages so that there's no error.

```go
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
```

- As we can see here, for sending messages, the ChatRoom only requires the name of the Participant as a string and it aptly handles the cases in which a person is no longer present in the chat room.
