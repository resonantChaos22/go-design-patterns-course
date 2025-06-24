`Observer` - is an object that wishes to be informed about events happening in the system. The entity generating the events will be `Observable`
- It is an intrusive approach as clients should have a way to subscribe.

## Simple Example

- Consider a function which when fired, will inform the doctors that you need help
```go
type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}
func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}
func (o *Observable) Fire(data any) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Person struct {
	Observable
	Name string
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

type Observer interface {
	Notify(data any)
}

type DoctorService struct{}

func (d *DoctorService) Notify(data any) {
	fmt.Printf("A doctor has been called for %s\n", data.(string))
}

```
- This is the definition of an `Observable` which fires an event to all its subscribers by calling their `Notify`
- `Observer`'s only job is to have a method to handle the data that is being sent

## Property Observer

- For monitoring individual properties, we can create a setter function and on fire we can fire a `PropertyChange` struct which can be consumed by the observers.
```go
type PropertyChange struct {
	Name  string
	Value any
}

type Person struct {
	Observable
	age  int
	name string
}

func (p *Person) Age() int {
	return p.age
}
func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	p.age = age
	p.Fire(PropertyChange{
		Name:  "Age",
		Value: p.age,
	})
}
```
- As you can see here, if there is a change in property "Age", the change is sent to all the subscribers.
```go
type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "Age" {
			if pc.Value.(int) >= 16 {
				fmt.Println("Congrats, you can drive now")
				t.o.Unsubscribe(t)
			} else {
				fmt.Println("We are monitoriing you")
			}
		}
	}
}

```
- And we handle it after verifying the parameter.
- Maybe create a common struct type to communicate information so that subscribers can work on it depending on whether they need to work on it

## Issue: Property Dependency

- There will be issues if we want to implement notifications for dependent properties as well.
- Let's say, we have a `CanVote()` function which returns whether we can vote or not based on age. We want there to be a notification when this changes as well.
- The only way to do this is to do this in the `SetAge()` function only
```go
func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{
		Name:  "Age",
		Value: p.age,
	})
	if p.CanVote() != oldCanVote {
		p.Fire(PropertyChange{
			Name:  "CanVote",
			Value: p.CanVote(),
		})
	}
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}
```

- As you can see here, if there are multiple dependencies, there will be problems in scaling and we will have to create a separate infrastructure.
