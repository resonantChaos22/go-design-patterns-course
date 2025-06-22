package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

// Generator
func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()

	return out
}

// Non-idiomatic approach, commonly used in cpp
type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{
		person:  person,
		current: -1,
	}
}
func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}
func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName

	default:
		panic("Not supported")
	}
}

func TestIteration() {
	p := Person{FirstName: "Alexander", MiddleName: "Grahamn", LastName: "Bell"}
	for _, name := range p.Names() {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	for name := range p.NamesGenerator() {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	for it := NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Printf("%s ", it.Value())
	}
	fmt.Println()

}
