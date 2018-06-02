package main

import (
	"os"
)

func main() {
	args := getArgs()
	all := parseArgs(args)

	img := openImage(os.Args[2])
	pixels := readPixels(img)
	x, y := readDimensions(img)
	modpxls := pxlToMod(pixels)
	execute(x, y, all, modpxls, os.Args[2])

}
