package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// VIOLATION OF LSP
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	return &Square{
		Rectangle: Rectangle{
			width:  size,
			height: size,
		},
	}
}

func (s *Square) SetHeight(size int) {
	s.width = size
	s.height = size
}

func (s *Square) SetWidth(size int) {
	s.width = size
	s.height = size
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println(expectedArea, actualArea)
}

func TestLSP() {
	rc := &Rectangle{
		width:  20,
		height: 20,
	}

	UseIt(rc)
	s := NewSquare(20)
	UseIt(s)
}
