- `Strategy Pattern` separates an algorithm into its 'skeleton' and concrete implementation steps, which can be varied at run-time
- This is used to abstract how a certain thing happens based on a strategy interface. We can change the strategy for the struct as we please.

## Example

```go
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

func (t *TextProcessor) AppendList(items []string) {
 s := t.listStrategy
 s.Start(&t.builder)
 for _, item := range items {
  s.AddListItem(&t.builder, item)
 }
 s.End(&t.builder)
}

tp := NewTextProcessor(Markdown)
tp.AppendList([]string{"foo", "bar", "baz"})
fmt.Println(tp)
tp.SetOutputFormat(Html)
tp.Reset()
tp.AppendList([]string{"foo", "bar", "baz"})
fmt.Println(tp)
```

- As you can see here, `TextProcessor` takes `ListStrategy` as an interface and defines its function on it on a higher order. It doesnt care how it is implemented.
