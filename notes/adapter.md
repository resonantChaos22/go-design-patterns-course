- `Adapter` is a construct which adapts an existing interface X to conform to required interface Y.
- For example, let's say you have the following interfaces.

```go
type Line struct {
 X1, Y1, X2, Y2 int
}
type VectorImage struct {
 Lines []Line
}
func NewRectangle(width, height int) *VectorImage {}

type Point struct {
 X, Y int
}
type RasterImage interface {
 GetPoints() []Point
}
func DrawPoints(owner RasterImage) string {}
```

- Now you have a `VectorImage` from `NewRectangle` and you want to draw it
- But you cant, as the `DrawPoints` function requires a `RasterImage` interface.
- So, for this to work, you need to create an `Adapter` which converts the `VectorImage` into `RasterImage`
- We can do it like this -

```go
// adapter
type vectorToRasterAdapter struct {
 points []Point
}

// implementing the interface
func (r *vectorToRasterAdapter) GetPoints() []Point {
 return r.points
}

//  function to convert data (Line -> Point)
func (r *vectorToRasterAdapter) addLine(line Line) {
 r.points = []
}

//  function to generate the second interface from first
func VectorToRaster(vi *VectorImage) RasterImage {
 adapter := vectorToRasterAdapter{}
 for _, line := range vi.Lines {
  adapter.addLine(line)
 }

 return &adapter
}

rc := NewRectangle(30, 10)
fmt.Println(DrawPoints(VectorToRaster(rc)))
```

- We can use the adapter function `VectorToRaster` to convert the `VectorImage` to a `RasterImage` so that we can use the `DrawPoints` function.
- This is how the `Adapter Pattern` helps us.

## Caching

- We can implement caching in go using `md5` and `json` packages.
- While creating the interface, if we dont want to convert the data of the first interface if we are passing an interface with the same data, we can convert the interface into a `[16]byte` and use it as a hash for the points generated.
- So, since we are hashing the entire object, if two objects are created with the exact same data, their hash will be the same.

```go
var pointCache map[[16]byte][]Point = map[[16]byte][]Point{}

func (r *vectorToRasterAdapter) addLineCached(line Line) {
 hash := func(obj interface{}) [16]byte {
  bytes, _ := json.Marshal(obj)
  return md5.Sum(bytes)
 }

 h := hash(line)
 if pts, ok := pointCache[h]; ok {
  r.points = append(r.points, pts...)
  return
 }

 // working on it if we dont find the hash
 r.points := []

 pointCache[h] = r.points
 fmt.Println("we have", len(r.points), "points")
}
```
