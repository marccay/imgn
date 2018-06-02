package main

import "fmt"

func main() {
	_, path, args := getArgs()
	all := parseArgs(args)

	fmt.Println(path)
	fmt.Println(all)
	/*
		img := openImage(path)
		pixels := readPixels(img)
		x, y := readDimensions(img)
		modpxls := pxlToMod(pixels)
		execute(x, y, all, modpxls, path)
	*/
}
