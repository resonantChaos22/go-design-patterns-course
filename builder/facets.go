package main

import "fmt"

type Person struct {
	StreetAddress string
	Postcode      string
	City          string

	Company      string
	Position     string
	AnnualIncome string
}

func (p *Person) Introduce() {
	fmt.Printf("I live at %s, %s, %s and I work at %s as a %s and I earn %s", p.StreetAddress, p.Postcode, p.City, p.Company, p.Position, p.AnnualIncome)
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{PersonBuilder: *b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{PersonBuilder: *b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// AddressBuilder methods
func (b *PersonAddressBuilder) At(streetAddr string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddr
	return b
}
func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}
func (b *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// JobBuilder methods
func (b *PersonJobBuilder) At(company string) *PersonJobBuilder {
	b.person.Company = company
	return b
}
func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}
func (b *PersonJobBuilder) Earning(annualIncome string) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func TestBuilderFacet() {
	pb := NewPersonBuilder()

	pb.
		Lives().
		At("212 Baker Street").
		In("London").
		WithPostCode("SW12BC").
		Works().
		At("Zapcom").
		AsA("Software Engineer").
		Earning("Nothing")

	p := pb.Build()
	p.Introduce()

}
