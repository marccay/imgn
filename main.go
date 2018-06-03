package main

import (
	"os"
)

func main() {
	args := getArgs()
	all := parseArgs(args)

	single(os.Args[2], all)
}

func single(file string, all multipleGroups) {
	img := openImage(file)
	pixels := readPixels(img)
	x, y := readDimensions(img)
	modpxls := pxlToMod(pixels)
	execute(x, y, all, modpxls, file)
}
