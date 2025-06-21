package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{
		Name:  "Circle",
		Color: color,
	}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{
		Name:  "Square",
		Color: color,
	}
}

func TestGeometric() {
	drawing := GraphicObject{
		Name: "My Drawing",
	}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Yellow"))

	group := GraphicObject{
		Name: "Group 1",
	}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
