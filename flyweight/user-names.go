package main

import (
	"fmt"
	"strings"
)

// unoptimized version
// over the course of large data, we might be storing many duplicate names
// at least duplicate first names and last names
type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{
		FullName: fullName,
	}
}

// flyweight optimized verion
// here, we are assuming that the total number of unique names wont exceed 256 (uint8)
var allNames []string

type User2 struct {
	names []uint8
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}

	return strings.Join(parts, " ")
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}

	return &result
}

func TestUserNames() {
	john := NewUser("John Doe")
	// jane := NewUser("Jane Doe")
	// alsoJane := NewUser("Jane Smith")
	fmt.Println(john.FullName)

	john2 := NewUser2("John Doe")
	// jane2 := NewUser2("Jane Doe")
	fmt.Println(john2.FullName())
}
