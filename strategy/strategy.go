package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (m *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (m *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (m *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("\t<li>" + item + "</li>\n")
}

type TextProcessor struct {
	builder      strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(fmt OutputFormat) *TextProcessor {
	tp := &TextProcessor{
		builder: strings.Builder{},
	}
	tp.SetOutputFormat(fmt)
	return tp
}
func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &MarkdownListStrategy{}

	case Html:
		t.listStrategy = &HtmlListStrategy{}

	default:
		panic("unimplemented output format")
	}
}
func (t *TextProcessor) AppendList(items []string) {
	s := t.listStrategy
	s.Start(&t.builder)
	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}
	s.End(&t.builder)
}
func (t *TextProcessor) Reset() {
	t.builder.Reset()
}
func (t *TextProcessor) String() string {
	return t.builder.String()
}

func TestStrategy() {
	tp := NewTextProcessor(Markdown)
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)
	tp.SetOutputFormat(Html)
	tp.Reset()
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

}
