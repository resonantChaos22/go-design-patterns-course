package main

import "fmt"

// developed and tested code
type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}
func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// new implementation of colored class
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func TestDecorator() {
	circle := Circle{Radius: 2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{
		Shape: &circle,
		Color: "Red",
	}
	fmt.Println(redCircle.Render())
	// redCircle.Resize() //	Resize method is not available as `ColoredShape` does not have the individual methods of the struct

	transparentRedCircle := TransparentShape{
		Shape:        &redCircle,
		Transparency: 0.93,
	}
	fmt.Println(transparentRedCircle.Render())
}
