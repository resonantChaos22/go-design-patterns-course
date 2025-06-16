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

// Does not follow OCP
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

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

//	Implementation of OCP following filter

type Specification interface {
	IsSatisfied(*Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec *ColorSpecification) IsSatisfied(product *Product) bool {
	return product.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec *SizeSpecification) IsSatisfied(product *Product) bool {
	return product.size == spec.size
}

type AndSpecification struct {
	specA Specification
	specB Specification
}

func (spec *AndSpecification) IsSatisfied(product *Product) bool {
	return spec.specA.IsSatisfied(product) && spec.specB.IsSatisfied(product)
}

type BetterFilter struct{}

func (bf BetterFilter) FilterValue(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
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
	for _, v := range f.FilterBySizeAndColor(products, large, green) {
		fmt.Println(v.name)
	}

	bf := BetterFilter{}
	SizeAndColorSpec := AndSpecification{
		specA: &ColorSpecification{color: green},
		specB: &SizeSpecification{size: large},
	}

	for _, v := range bf.FilterValue(products, &SizeAndColorSpec) {
		fmt.Println(v.name)
	}

}
