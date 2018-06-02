package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func openImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		fmt.Println("issue opening file, assure file is a jpeg")
		os.Exit(1)
	}

	return img
}

func createDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				os.Exit(1)
			}
		} else {
			os.Exit(1)
		}
	}
}

func writeToFile(img image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	err = jpeg.Encode(w, img, nil)
	w.Flush()
}
