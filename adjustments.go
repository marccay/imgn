package main

import (
	"fmt"
	"os"
)

func (pixels allModifiedrgba) brightness(adjustment float64) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))
	// -1 to 1 value accepted
	if adjustment > 1 || adjustment < -1 {
		fmt.Println("brightness adjustment value is out of bounds, choose float between -1 and 1")
		os.Exit(1)
	}

	// shade(darker) adjustment float64 more than -1
	// smaller number == darker tint, so adjusted for
	// that
	if adjustment < 0 {
		adj := 1 - (-adjustment)
		for n, pxl := range pixels {
			new[n].r, new[n].g, new[n].b = pxl.shade(adj)
			new[n].a = pxl.a
		}
		// tint (brighter) adjustment float64 less than 1
	} else if adjustment > 0 {
		adj := adjustment
		for n, pxl := range pixels {
			new[n].r, new[n].g, new[n].b = pxl.tint(adj)
			new[n].a = pxl.a
		}
	}

	return new
}

func (pixel modifiedrgba) shade(adj float64) (r uint8, g uint8, b uint8) {
	r = uint8(float64(pixel.r) * adj)
	g = uint8(float64(pixel.g) * adj)
	b = uint8(float64(pixel.b) * adj)
	return r, g, b
}

func (pixel modifiedrgba) tint(adj float64) (r uint8, g uint8, b uint8) {
	r = pixel.r + uint8(float64(255-pixel.r)*adj)
	g = pixel.g + uint8(float64(255-pixel.g)*adj)
	b = pixel.b + uint8(float64(255-pixel.b)*adj)
	return r, g, b
}

func (pixels allModifiedrgba) contrast(adjustment float64) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))

	if adjustment < 0.0 || adjustment > 3.0 {
		fmt.Println("contrast adjustment value is out of bounds, choose float between 0 and 3.0")
		os.Exit(1)
	}

	// w/o factor negative is between 1 and 0, close to 0 reduce contrast
	// reasonable for positive is between 1 and 3, higher num more contrast
	for n, pxl := range pixels {
		newr := ((adjustment * (float64(pxl.r) - 128.0)) + 128.0)
		new[n].r = toUint8(newr)
		newg := ((adjustment * (float64(pxl.g) - 128.0)) + 128.0)
		new[n].g = toUint8(newg)
		newb := ((adjustment * (float64(pxl.b) - 128.0)) + 128.0)
		new[n].b = toUint8(newb)
		new[n].a = pxl.a
	}
	return new
}

func (pixels allModifiedrgba) desaturate(shade string, custom desatFormula) allModifiedrgba {
	new := make([]modifiedrgba, len(pixels))
	//i := 0
	for n, pxl := range pixels {
		var gray uint8
		if shade == "luminosity" {
			// luminosity formula
			lumi := 0.21*float64(pxl.r) + 0.72*float64(pxl.g) + 0.07*float64(pxl.b)
			gray = uint8(lumi)
		} else if shade == "average" {
			avg := (float64(pxl.r) + float64(pxl.g) + float64(pxl.b)) / 3
			gray = uint8(avg)
		} else if shade == "lightness" {
			//lightness : average between max - min values
			max := pxl.max()
			min := pxl.min()
			light := (float64(max) + float64(min)) / 2
			gray = uint8(light)
		} else if shade == "custom" {
			r := custom.r
			g := custom.g
			b := custom.b
			cus := r*float64(pxl.r) + g*float64(pxl.g) + b*float64(pxl.b)
			gray = uint8(cus)
		}
		new[n].r = gray
		new[n].g = gray
		new[n].b = gray
		new[n].a = pxl.a
	}

	return new
}

func (pixels allModifiedrgba) highlights(adjustment float64) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))

	for n, pxl := range pixels {
		new[n] = pxl.modd(0, 280)
	}
	return new
}

func (pixels allModifiedrgba) opacity(alpha uint8) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))

	for n, pxl := range pixels {
		new[n] = pxl
		new[n].a = alpha
	}
	return new
}
