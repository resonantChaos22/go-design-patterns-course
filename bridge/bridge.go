package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(size int)
}

// vector renderer implementation
type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing through vector a circle of radius ", radius)
}
func (v *VectorRenderer) RenderSquare(size int) {
	fmt.Println("Drawing through vector a square of size ", size)
}

// raster render implementation
type RastererRenderer struct {
	Dpi int
}

func (r *RastererRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing through raster pixels for a circle of radius ", radius)
}
func (r *RastererRenderer) RenderSquare(size int) {
	fmt.Println("Drawing through raster a square of size ", size)
}

// circle implementation
type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{
		renderer: renderer,
		radius:   radius,
	}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}
func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Square struct {
	renderer Renderer
	size     int
}

func NewSquare(renderer Renderer, size int) *Square {
	return &Square{
		renderer: renderer,
		size:     size,
	}
}
func (s *Square) Draw() {
	s.renderer.RenderSquare(s.size)
}

func TestBridge() {
	raster := RastererRenderer{Dpi: 10}
	vector := VectorRenderer{}

	circle := NewCircle(&raster, 4)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()

	vCircle := NewCircle(&vector, 7)
	vCircle.Draw()

	rSquare := NewSquare(&raster, 8)
	vSquare := NewSquare(&vector, 3)

	rSquare.Draw()
	vSquare.Draw()
}
