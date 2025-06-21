- We want to augment an object with Additional Functionality so as to not break Open Close Functionality and Single Responsibility Principle
- `Decorator` facilitates the addition of behaviors to individual objects through embedding.

#### Problem in Golang

- In Golang, if a struct is composited with two structs with the same field names, there will be a consistency problem as we need to set both the ages individually and there could be cases in which we can set the two ages differently which would cause the whole thing to break.

```go
type Bird struct {
 Age int
}
type Lizard struct {
 Age int
}
type Dragon struct {
 Bird
 Lizard
}

d := Dragon{}
//d.Age = 10   -> Cant do this as both parent structs have the same field
d.Bird.Age = 10
d.Lizard.Age = 10
//  This could lead to consistency problem
```

- A better way to do this would be to create an interface of all the beings with the getter and setter methods for the same and making sure that the `age` is a private field so that it can not be affected directly by any functions utilizing this.

```go
type Aged interface {
 Age() int
 SetAge(age int)
}

type Dragon struct {
 bird   Bird
 lizard Lizard
}
func (d *Dragon) Age() int {
 return d.bird.Age()
}
func (d *Dragon) SetAge(age int) {
 d.bird.SetAge(age)
 d.lizard.SetAge(age)
}
func (d *Dragon) Fly() {
 d.bird.Fly()
}
func (d *Dragon) Crawl() {
 d.lizard.Crawl()
}
```

- Then we implement all the functions directly. The main reason to do this is so that we dont directly change the age of either bird or lizard by mistake and we only use the given functions to us to interact with the``Dragon

## Decorator

- Here's an example of Decorator at work

```go
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

type ColoredShape struct {
 Shape Shape
 Color string
}

func (c *ColoredShape) Render() string {
 return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

redCircle := ColoredShape{
 Shape: &Circle{
  Radius: 2
 },
 Color: "Red"
}
```

- The `ColoredShape` is the decorator at work and works for any struct that implements the `Shape` interface, for example `Square
- The issue here is that the object of `ColoredShape` loses access to the individual methods of `Circle`, so the `redCicle` has no way to utilize the `Resize` function.
- The advantage is that `Decorator`s can be composed

```go
type TransparentShape struct {
 Shape        Shape
 Transparency float32
}

func (t *TransparentShape) Render() string {
 return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

trc := TransparentShape{
 Shape:        &redCircle,
 Transparency: 0.93,
}
```

- Now `trc` will have both the fields
