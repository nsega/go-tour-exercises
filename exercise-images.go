package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (a *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (a *Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{200, 100}}
}

func (a *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x % 256), uint8(y % 256), uint8((x * y) % 256), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(&m)
}
