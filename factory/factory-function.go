package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func TestFactoryFunc() {
	//	Without Factory Function
	p1 := Person{
		Name:     "Shreyash",
		Age:      24,
		EyeCount: 2,
	}

	//	With Factory Function
	p2 := NewPerson("Shreyash", 24)

	fmt.Println(p1)
	fmt.Println(p2)
}

//	Note that with factory function, we can do validation and fill some default values or values
//	that can be generated with some values so that part of code is not repeated every time.
