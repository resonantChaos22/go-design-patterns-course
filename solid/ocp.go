package main

import "fmt"

type Color int

const (
	red Color = iota
	blue
	green
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func TestOCP() {
	apple := Product{name: "Apple", color: green, size: small}
	tree := Product{name: "Tree", color: green, size: large}
	house := Product{name: "House", color: blue, size: large}

	products := []Product{apple, tree, house}
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Println(v.name)
	}

}
