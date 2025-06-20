package main

import "fmt"

type Employee2 struct {
	Name         string
	Position     string
	AnnualIncome int
}

type role int

const (
	Developer role = iota
	Manager
)

func NewEmployee2(role role) *Employee {
	switch role {
	case Developer:
		return &Employee{Postion: "Developer", AnnualIncome: 60000}
	case Manager:
		return &Employee{Postion: "Manager", AnnualIncome: 80000}
	default:
		panic("No such role")
	}
}

func TestPrototypeFactory() {
	m := NewEmployee2(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
