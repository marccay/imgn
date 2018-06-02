package main

import (
	"image"
	"image/color"
)

func modify(x int, y int, pixels allModifiedrgba) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, x, y))

	bounds := img.Bounds()
	i := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			img.SetRGBA(x, y, color.RGBA{pixels[i].r, pixels[i].g, pixels[i].b, pixels[i].a})
			i++
		}
	}
	return img
}

func pxlToMod(ogpixels []originalrgba) allModifiedrgba {
	modpixels := make(allModifiedrgba, len(ogpixels))
	for i, x := range ogpixels {
		modpixels[i].r = uint8(x.r / 0x101)
		modpixels[i].g = uint8(x.g / 0x101)
		modpixels[i].b = uint8(x.b / 0x101)
		modpixels[i].a = uint8(x.a / 0x101)
	}
	return modpixels
}
