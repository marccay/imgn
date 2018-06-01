package main

import (
	"image"
)

type originalrgba struct {
	r, g, b, a uint32
}
type modifiedrgba struct {
	r, g, b, a uint8
}
type allModifiedrgba []modifiedrgba
type desatFormula struct {
	r, g, b float64
}

// read individual rgba values from each picture, consecutively inputing
// value to originalrgba as return in uint32
func readPixels(img image.Image) []originalrgba {
	bounds := img.Bounds()
	pixels := make([]originalrgba, bounds.Dx()*bounds.Dy())

	i := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[i].r = r
			pixels[i].b = b
			pixels[i].g = g
			pixels[i].a = a
			i++
		}
	}
	return pixels
}

func readDimensions(img image.Image) (x int, y int) {
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy()
}
