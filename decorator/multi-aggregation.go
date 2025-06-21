package main

import "fmt"

// Incorrect way of aggregation
type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	Bird
	Lizard
}

func (d *Dragon) GetAge() int {
	return d.Bird.Age
}
func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

// Correct Way of Aggregation using Decorator
type Aged interface {
	Age() int
	SetAge(age int)
}
type NewBird struct {
	age int
}

func (b *NewBird) Age() int {
	return b.age
}
func (b *NewBird) SetAge(age int) {
	b.age = age
}
func (b *NewBird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

type NewLizard struct {
	age int
}

func (l *NewLizard) Age() int {
	return l.age
}
func (l *NewLizard) SetAge(age int) {
	l.age = age
}
func (l *NewLizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

type NewDragon struct {
	bird   NewBird
	lizard NewLizard
}

func (d *NewDragon) Age() int {
	return d.bird.Age()
}
func (d *NewDragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}
func (d *NewDragon) Fly() {
	d.bird.Fly()
}
func (d *NewDragon) Crawl() {
	d.lizard.Crawl()
}

func BirthNewDragon() *NewDragon {
	return &NewDragon{
		bird:   NewBird{},
		lizard: NewLizard{},
	}
}

func TestMultiAggregation() {
	d := Dragon{}
	// d.Age = 10	//	doesnt work as both objects have individual fields called Age
	d.SetAge(10)
	d.Fly()
	d.Crawl()

	fmt.Println()
	nd := BirthNewDragon()
	nd.SetAge(9)
	nd.Fly()
	nd.Crawl()
}
