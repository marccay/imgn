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