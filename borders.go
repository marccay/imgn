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
