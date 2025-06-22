package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{
		filename: filename,
	}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}

	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{
		filename: filename,
	}
}

func TestVirtualProxy() {
	bmp := NewBitmap("demo.png")
	DrawImage(bmp)

	//	we are still loading the image even though we are not using the image
	fmt.Println()
	_ = NewBitmap("asta.png")

	fmt.Println()
	lbmp := NewLazyBitmap("sapphire.png")
	DrawImage(lbmp)

	fmt.Println()
	_ = NewLazyBitmap("someday.png")
}
