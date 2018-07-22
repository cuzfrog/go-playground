package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	data image.Rectangle
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.data
}

func (img Image) At(x, y int) color.Color {
	return img.data.At(x, y)
}

func main() {
	m := Image{image.Rect(0, 0, 10, 10)}
	pic.ShowImage(m)
}
