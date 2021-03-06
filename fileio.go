package main

import (
	"bufio"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func openImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("issue opening image file\n%v\n", err)
	}
	defer f.Close()

	var img image.Image
	img, err = jpeg.Decode(f)
	if err != nil {
		// try with png.decode if not jpeg
		img, err = png.Decode(f)
		if err != nil {
			log.Fatalf("issue opening file, assure file is a jpeg or png\n%v\n", err)
		}
	}

	return img
}

func createDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				log.Fatalf("error making output dir\n%v\n", err)
			}
		} else {
			os.Exit(1)
		}
	}
}

func writeToFile(img image.Image, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("error writing image.Image to file\n%v\n", err)
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
	sclock := strings.Split(cclock[0], ":")
	clock := strings.Join(sclock, "-")
	return date + "-" + clock
}
