package main

import (
	"log"
)

func (pixels allModifiedrgba) brightness(adjustment float64) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))
	// -1 to 1 value accepted
	if adjustment > 1 || adjustment < -1 {
		log.Fatalf("brightness adjustment value %v is out of bounds, choose float between -1 and 1\n", adjustment)
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

	var normalized float64
	normalized = adjustment + 1.0

	if adjustment < -1.0 || adjustment > 3.0 {
		log.Fatalf("contrast adjustment value %v is out of bounds, choose float between -1.0 and 2.0\n", adjustment)
	}

	// w/o factor negative is between 1 and 0, close to 0 reduce contrast
	// reasonable for positive is between 1 and 3, higher num more contrast
	for n, pxl := range pixels {
		newr := ((normalized * (float64(pxl.r) - 128.0)) + 128.0)
		new[n].r = toUint8(newr)
		newg := ((normalized * (float64(pxl.g) - 128.0)) + 128.0)
		new[n].g = toUint8(newg)
		newb := ((normalized * (float64(pxl.b) - 128.0)) + 128.0)
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
			r, g, b := custom.r, custom.g, custom.b
			cus := r*float64(pxl.r) + g*float64(pxl.g) + b*float64(pxl.b)
			gray = uint8(cus)
		}
		new[n].r, new[n].g, new[n].b = gray, gray, gray
		new[n].a = pxl.a
	}

	return new
}

func (pixels allModifiedrgba) highlights(adjustment float64) allModifiedrgba {
	if adjustment > 1.0 || adjustment < 0.0 {
		log.Fatalf("highlights recovery value %v is out of bounds. Choose a number between 0 and 1", adjustment)
	}

	new := make(allModifiedrgba, len(pixels))
	b := uint8(240)
	for n, pxl := range pixels {
		if pxl.r > b && pxl.g > b && pxl.b > b {
			new[n].r = reduceHighlights(pxl.r, adjustment)
			new[n].g = reduceHighlights(pxl.g, adjustment)
			new[n].b = reduceHighlights(pxl.b, adjustment)
		} else {
			new[n].r, new[n].g, new[n].b = pxl.r, pxl.g, pxl.b
		}
		new[n].a = pxl.a
	}
	return new
}

func reduceHighlights(color uint8, adjustment float64) uint8 {
	b := 240.0
	delta := (float64(color) - b) / (100.0 / (255.0 - b))
	new := float64(color) - (delta * (255.0 - b) * adjustment)
	return uint8(new)
}

func (pixels allModifiedrgba) shadows(adjustment float64) allModifiedrgba {
	if adjustment > 1.0 || adjustment < 0.0 {
		log.Fatalf("shadow recovery value %v is out of bounds. Choose a number between 0 and 1", adjustment)
	}
	new := make(allModifiedrgba, len(pixels))
	b := uint8(20)
	for n, pxl := range pixels {
		if pxl.r < b && pxl.g < b && pxl.b < b {
			new[n].r = reduceShadows(pxl.r, adjustment)
			new[n].g = reduceShadows(pxl.g, adjustment)
			new[n].b = reduceShadows(pxl.b, adjustment)
		} else {
			new[n].r, new[n].g, new[n].b = pxl.r, pxl.g, pxl.b
		}
		new[n].a = pxl.a
	}
	return new
}

func reduceShadows(color uint8, adjustment float64) uint8 {
	b := 20.0
	delta := (float64(color) - b) / (-b)
	new := float64(color) + (delta * (b) * adjustment)
	return uint8(new)
}

func (pixels allModifiedrgba) opacity(alpha uint8) allModifiedrgba {
	new := make(allModifiedrgba, len(pixels))

	for n, pxl := range pixels {
		new[n] = pxl
		new[n].a = alpha
	}
	return new
}
