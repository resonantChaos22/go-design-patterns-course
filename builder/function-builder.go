package main

import "fmt"

type PersonA struct {
	name    string
	postion string
}

type personAMod func(*PersonA)
type PersonABuilder struct {
	actions []personAMod
}

func (b *PersonABuilder) Called(name string) *PersonABuilder {
	b.actions = append(b.actions, func(pa *PersonA) {
		pa.name = name
	})

	return b
}

func (b *PersonABuilder) Is(positon string) *PersonABuilder {
	b.actions = append(b.actions, func(pa *PersonA) {
		pa.postion = positon
	})

	return b
}

func (b *PersonABuilder) Build() *PersonA {
	p := PersonA{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func introducePerson(person *PersonA) {
	fmt.Printf("Hi, My name is %s, and I work as a %s", person.name, person.postion)
}

type personBuild func(*PersonABuilder)

func Introduce(action personBuild) {
	pb := PersonABuilder{}
	action(&pb)
	//	at this point, all the actions required to create a function have been added
	introducePerson(pb.Build())
}

func TestFunctionalBuilder() {
	b := PersonABuilder{}
	p := b.Called("Dmitri").Is("Technician").Build()
	fmt.Println(p.name, p.postion)

	Introduce(func(pa *PersonABuilder) {
		pa.Called("Shreyash").Is("Developer")
	})
}

//	TODO try to use build params here
