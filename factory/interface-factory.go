package main

import "fmt"

type PersonA interface {
	SayHello()
}

type personA struct {
	name string
	age  int
}

func (p *personA) SayHello() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", p.name, p.age)
}

func NewPersonA(name string, age int) PersonA {
	return &personA{
		name: name,
		age:  age,
	}
}

func TestInterfacFactory() {
	newPersonA := NewPersonA("Shreyash", 24)
	newPersonA.SayHello()
}
