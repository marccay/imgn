package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := getArgs()
	all := parseArgs(args)

	path, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Println("abs failed")
		os.Exit(1)
	}

	single(path, all)
}

func single(file string, all multipleGroups) {
	img := openImage(file)
	fmt.Println("image opened")
	pixels := readPixels(img)
	x, y := readDimensions(img)
	modpxls := pxlToMod(pixels)
	execute(x, y, all, modpxls, file)
}
