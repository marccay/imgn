package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"
)

func openImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("issue opening image file")
		os.Exit(1)
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
				fmt.Println("error making output dir")
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

func stringTime() string {
	ttime := time.Now().String()
	stime := strings.Split(ttime, " ")
	date := stime[0]
	cclock := strings.Split(stime[1], ".")
	clock := cclock[0]
	return date + "-" + clock
}
