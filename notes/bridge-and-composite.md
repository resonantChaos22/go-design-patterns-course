- `Bridge Pattern` prevents a `Cartesian Product` complexity explosion
- Take an example - If we are building a Thread Scheduler which are of two types - Preemptive and Cooperating and we are building them for Windows and Unix, we will have to make 7 different structs for it.
- ![[Bridge Pattern 2025-06-21 16.43.07.excalidraw|2000x100]]
- `Bridge` is a mechanism that decouple an interface from an implementation

## Simple Bridge usage

- Take the case that we have different shapes and different renderers to draw that shape. We can have shapes like Square, Circle, etc and we can have different renderers like Vector Renderer and Circle Renderer.
- So normally, we would have to create `CircleVectorRenderer`, `CircleRasterRenderer` and so on and this would lead to so many structs.
- Instead, we can define an interface `Renderer` which needs to have functions of how to actually draw the particular shape.
- And then we can pass the `Renderer` inside the different types of objects so that they can utilize the render function of either of the two renderers

```go
type Renderer interface {
 RenderCircle(radius float32)
}

type RastererRenderer struct {
 Dpi int
}

func (r *RastererRenderer) RenderCircle(radius float32) {
 fmt.Println("Drawing through raster pixels for a circle of radius ", radius)
}

type Circle struct {
 renderer Renderer
 radius   float32
}

func (c *Circle) Draw() {
 c.renderer.RenderCircle(c.radius)
}

raster := RastererRenderer{Dpi: 10}
circle := NewCircle(&raster, 4)
circle.Draw()
```

## Composite Pattern

- Composition lets us make compound objects
- Objects use other objects fields/methods via embedding
- Composite Design pattern is used to treat both single and composite objects uniformly

```go
type GraphicObject struct {
 Name, Color string
 Children    []GraphicObject
}
```

- The methods for this object should be such that they work correctly irrespective of whether the `Children` field is empty or not. Something like `print` ->

```go
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
```

- As you can see here, this function is called recursively to print all the children in case the object does have children.

#### Neuron Network

- Another example of this is a `Neural Network` where there are `Neurons` and `Neuron Layers` and there should be a `Connect` function which can connect both Neuron to Neuron Layers.
- Doing this via `Composite Pattern` would make us create an interface which can basically give us a function to get all the neurons in a single neuron or a neuron layer so that we can connect them.

```go
type NeuronInterface interface {
 Iter() []*Neuron
}

type Neuron struct {
 In, Out []*Neuron
}
func (n *Neuron) Iter() []*Neuron {
 return []*Neuron{n}
}


type NeuronLayer struct {
 Neurons []Neuron
}
func (n *NeuronLayer) Iter() []*Neuron {
 result := make([]*Neuron, 0)
 for _, neuron := range n.Neurons {
  result = append(result, &neuron)
 }

 return result
}

func Connect(left, right NeuronInterface) {
 for _, l := range left.Iter() {
  for _, r := range right.Iter() {
   l.ConnectTo(r)
  }
 }
}
```

- The `Connect` function works on the interface due to which we can treat a single `Neuron` and a combination of `Neuron`s similarly
