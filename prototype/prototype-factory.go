package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type OfficeAddress struct {
	Suite         int
	StreetAddress string
	City          string
}

type Employee struct {
	Name   string
	Office OfficeAddress
}

type office int

const (
	MainOffice office = iota
	AuxOffice
)

func (em *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(&em)

	result := Employee{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&result)

	return &result
}

var mainOfficeProto = Employee{
	Office: OfficeAddress{
		StreetAddress: "123 East Drive",
		City:          "London",
	},
}
var auxOfficeProto = Employee{
	Office: OfficeAddress{
		StreetAddress: "456 West Drive",
		City:          "London",
	},
}

func newEmployeeFromProto(proto *Employee, name string, suite int) *Employee {
	newEmployee := proto.DeepCopy()
	newEmployee.Name = name
	newEmployee.Office.Suite = suite

	return newEmployee
}

func NewEmployee(office office, name string, suite int) *Employee {
	switch office {
	case MainOffice:
		return newEmployeeFromProto(&mainOfficeProto, name, suite)

	case AuxOffice:
		return newEmployeeFromProto(&auxOfficeProto, name, suite)

	default:
		panic("No such office")
	}
}

func TestPrototypeFactory() {
	john := NewEmployee(MainOffice, "John", 102)
	mark := NewEmployee(AuxOffice, "Mark", 205)

	fmt.Println(john.Name, john.Office)
	fmt.Println(mark.Name, mark.Office)
}
