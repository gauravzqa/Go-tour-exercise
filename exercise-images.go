package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	h, w  int
	color uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.h, i.w)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}

func main() {
	m := Image{200, 200, 200}
	pic.ShowImage(m)
}
