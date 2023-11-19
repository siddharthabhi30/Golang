package main

import (
	"image"
	"image/color"
)

type Image struct {
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Max: image.Point{
			X: 10,
			Y: 10,
		},
		Min: image.Point{
			X: 0,
			Y: 0,
		},
	}
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
}
