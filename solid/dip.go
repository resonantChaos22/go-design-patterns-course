package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module (basically just data)
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high level module (functions on data) || INCORRECT IMPLEMENTATION ACCORDING TO DIP
// type Research struct {
// 	relationships Relationships
// }

// func (r *Research) Investigate() {
// 	relations := r.relationships.relations
// 	for _, rel := range relations {
// 		if rel.relationship == Parent && rel.from.name == "John" {
// 			fmt.Printf("%s has a child called %s\n", rel.from.name, rel.to.name)
// 		}
// 	}
// }

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	children := make([]*Person, 0)

	for _, rel := range r.relations {
		if rel.relationship == Parent && rel.from.name == name {
			children = append(children, rel.to)
		}
	}

	return children
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, child := range r.browser.FindAllChildrenOf("John") {
		fmt.Printf("Johb has a child called %s\n", child.name)
	}
}

func TestDIP() {
	rels := Relationships{}
	rels.AddParentAndChild(&Person{name: "John"}, &Person{name: "Mark"})
	rels.AddParentAndChild(&Person{name: "Mark"}, &Person{name: "Chris"})

	r := Research{
		browser: &rels,
	}
	r.Investigate()
}
