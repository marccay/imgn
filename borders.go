<<<<<<< HEAD
package main

func (pixels allModifiedrgba) border(x int, y int, borderWidth int) allModifiedrgba {
	new := make(allModifiedrgba, (x+borderWidth)*(y+borderWidth))

	for x, pxl := range pixels {
		if pxl % x < y 
	}

	return new
}

func (pixels allModifiedrgba) flip(x int, y int) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))

	d := (new % x) * y 

	return new

}
=======
package main

import (
	"log"
)

func (pixels allModifiedrgba) border(x int, y int, borderWidth int, borderColor modifiedrgba) (int, int, allModifiedrgba) {
	if borderWidth < 0 {
		log.Fatalf("set border size greater than 0")
	}

	newX := x + borderWidth + borderWidth
	newY := y + borderWidth + borderWidth

	new := make(allModifiedrgba, newX*newY)

	pixelsCount := 0
	n := 0

	for xx := 0; xx < newX; xx++ {
		for yy := 0; yy < newY; yy++ {
			if xx < borderWidth || xx >= borderWidth+x {
				new[n].r, new[n].g, new[n].b, new[n].a = borderColor.r, borderColor.g, borderColor.b, borderColor.a
			} else {
				if yy < borderWidth || yy >= borderWidth+y {
					new[n].r, new[n].g, new[n].b, new[n].a = borderColor.r, borderColor.g, borderColor.b, borderColor.a
				} else {
					new[n].r, new[n].g, new[n].b, new[n].a = pixels[pixelsCount].r, pixels[pixelsCount].g, pixels[pixelsCount].b, pixels[pixelsCount].a
					pixelsCount++
				}
			}
			n++
		}
	}

	return newX, newY, new
}

func (pixels allModifiedrgba) borderPercent(x int, y int, borderPercentage float64, borderColor modifiedrgba) (int, int, allModifiedrgba) {
	borderWidth := int(borderPercentage * float64(y))
	return pixels.border(x, y, borderWidth, borderColor)
}
>>>>>>> 36b33bb9b04fb8ff5db5de9c30e963bd7dbbd863
