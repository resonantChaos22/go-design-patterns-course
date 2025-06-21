package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}

//	given interface

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		Lines: []Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// the interface that we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()

	//	setting maxX and maxY
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}

		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	maxX += 1
	maxY += 1

	//	pre-allocate
	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	//	setting the points
	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	//	creating string
	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// adapter
type vectorToRasterAdapter struct {
	points []Point
}

// implementing the interface
func (r *vectorToRasterAdapter) GetPoints() []Point {
	return r.points
}

// conversion of data, we are using the cached version
func (r *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			r.points = append(r.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			r.points = append(r.points, Point{x, top})
		}
	}

	fmt.Println("we have", len(r.points), "points")
}

// caching the adapter generated points so we dont have to do it again and again
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

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)

	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			r.points = append(r.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			r.points = append(r.points, Point{x, top})
		}
	}

	pointCache[h] = r.points
	fmt.Println("we have", len(r.points), "points")
}

// converting the first interface into the required interface
func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}

	return &adapter
}

func TestAdapter() {
	//	first interface
	rc := NewRectangle(30, 10)
	rc2 := NewRectangle(30, 10)

	//	adapter
	a := VectorToRaster(rc)

	//	requires second interface
	fmt.Println(DrawPoints(a))
	fmt.Println(DrawPoints(VectorToRaster(rc2)))

	//	recomputation happens as the `VectorImage` is different
	fmt.Println(DrawPoints(VectorToRaster(NewRectangle(30, 20))))
}
