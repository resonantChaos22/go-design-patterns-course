package main

import (
	"container/list"
	"fmt"
)

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

type PropertyChange struct {
	Name  string
	Value any
}

type Person struct {
	Observable
	age  int
	name string
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

func (p *Person) Age() int {
	return p.age
}
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

func NewPerson(name string, age int) *Person {
	return &Person{
		Observable: Observable{
			subs: new(list.List),
		},
		name: name,
		age:  age,
	}
}

type Observer interface {
	Notify(data any)
}

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

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data any) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func TestObserverPattern() {
	p := NewPerson("Sabu", 13)
	t := &TrafficManagement{
		o: p.Observable,
	}
	e := &ElectoralRoll{}
	p.Subscribe(t)
	p.Subscribe(e)

	for i := 14; i <= 20; i++ {
		fmt.Println("Setting the age to ", i)
		p.SetAge(i)
	}

}
