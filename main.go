package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var bunchFiles sync.WaitGroup

func main() {
	args := getArgs()
	all := parseArgs(args)

	path, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Println("abs failed")
		os.Exit(1)
	}

	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println("trouble accessing path info")
		os.Exit(1)
	}

	if stat.IsDir() {
		multiplex(path, all)
	} else {
		omniplex(path, all)
	}

}

func omniplex(file string, all multipleGroups) {
	defer bunchFiles.Done()
	img := openImage(file)
	pixels := readPixels(img)
	x, y := readDimensions(img)
	modpxls := pxlToMod(pixels)
	execute(x, y, all, modpxls, file)
}

func multiplex(path string, all multipleGroups) {
	var bunchFilepaths []string
	err := filepath.Walk(path, func(fp string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("failed to access file path")
			os.Exit(2)
		}
		if info.IsDir() {
			return nil
		}
		bunchFilepaths = append(bunchFilepaths, fp)
		return nil
	})
	if err != nil {
		fmt.Println("error walking the path")
		os.Exit(2)
	}
	bunchFiles.Add(len(bunchFilepaths))
	for _, f := range bunchFilepaths {
		go omniplex(f, all)
	}
	bunchFiles.Wait()
}
