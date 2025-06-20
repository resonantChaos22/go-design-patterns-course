package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress string
	City          string
	Country       string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)

	return &q
}

func (p *Person) DeepCopyThroughSerialization() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func TestDeepCopy() {
	john := Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Jane", "Mark", "Adam"},
	}
	//	wrong, copies the pointer, so both john and jane share the same address on memory
	// jane := john
	// jane.Name = "Jane"
	// jane.Address.StreetAddress = "212 Baker Street"

	//	how to modiy without any unexpected problem
	jane := john
	jane.Name = "Jane"
	jane.Address = &Address{
		StreetAddress: "212 Baker Street",
		City:          "London",
		Country:       "UK",
	}

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	//	using the copy method
	fmt.Println()
	mark := john.DeepCopy()
	mark.Name = "Mark"
	mark.Address.StreetAddress = "414 Capybara Street"
	mark.Friends = append(mark.Friends, "Angela")

	fmt.Println(john.Name, john.Address, john.Friends)
	fmt.Println(mark.Name, mark.Address, mark.Friends)

	//	deep copy through serialization
	fmt.Println()
	adam := john.DeepCopyThroughSerialization()
	adam.Name = "Adam"
	adam.Address.StreetAddress = "Garden Of Eve"
	adam.Friends = append(adam.Friends, "Eve")

	fmt.Println(john.Name, john.Address, john.Friends)
	fmt.Println(adam.Name, adam.Address, adam.Friends)
}
