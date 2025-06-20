package main

import (
	"fmt"
	"strings"
)

const (
	IndentSize = 2
)

type HtmlElement struct {
	name     string
	text     string
	elements []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", IndentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", IndentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)

	return b
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func TestStringBuilder() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()

	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	//	Builder Pattern
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	//	Fluent one
	fb := NewHtmlBuilder("ul")
	fb.AddChildFluent("li", "item 1").AddChildFluent("li", "item 2")

	fmt.Println(fb.String())

}
