package main

import "fmt"

type Employee struct {
	Name         string
	Postion      string
	AnnualIncome int
}

// functional approach to generate different types of employee
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Postion:      position,
			AnnualIncome: annualIncome,
		}
	}
}

// structural Approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactoryB(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{Name: name, Postion: f.Position, AnnualIncome: f.AnnualIncome}
}

func TestFactoryGenerator() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	devA := developerFactory("Shreyash")
	manA := managerFactory("Adam")

	fmt.Println(devA)
	fmt.Println(manA)

	devFactory := NewEmployeeFactoryB("Developer", 60000)
	qaFactory := NewEmployeeFactoryB("QA Engineer", 40000)

	devB := devFactory.Create("Shreyash")
	qaB := qaFactory.Create("Christy")

	fmt.Println(devB)
	fmt.Println(qaB)
}
