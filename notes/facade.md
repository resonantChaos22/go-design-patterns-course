- `Facade` design pattern is used to provide a simple, easy to understand user interface over a large and sophisticated body of code.
- Consider an example that you want to create a system to show what is on a terminal console.
- A terminal console has `Viewport`s which have their individual `Buffer` and the `Buffer` contains all the text that is shown on the terminal.

```go
type Buffer struct {
 width, height int
 buffer        []rune
}
type Viewport struct {
 buffer *Buffer
 offset int
}

type Console struct {
 buffers   []*Buffer
 viewports []*Viewport
 offset    int
}
func (c *Console) GetCharacterAt(index int) rune {
 return c.viewports[0].GetCharacterAt(index)
}

// default console has 1 buffer and 1 viewport.
func NewDefaultConsole() *Console {
 b := NewBuffer(200, 150)
 v := NewViewport(b)

 return &Console{
  buffers:   []*Buffer{b},
  viewports: []*Viewport{v},
  offset:    0,
 }
}
c := NewDefaultConsole()
c.GetCharacterAt(10)
```

- Now, both `Buffer` and `Viewport` have individual methods to get the character at a particular position but we provide a `NewDefaultConsole` which works as a `Facade` to make it simpler for you to get the character without knowing the internal workings of buffer and viewport.
- It's basically an example of `Encapsulation`
- The users have the choice to use the advanced options and directly manipulate Buffer and Viewport but they also have the option to simply work on the `Console` object and leave all the complexities of how `Viewport` and `Buffer` work alone.
