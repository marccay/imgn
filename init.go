package main

import (
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func execute(x int, y int, groups multipleGroups, pixels allModifiedrgba, path string) {
	fullpath, _ := filepath.Abs(path)
	dir := filepath.Dir(fullpath)
	ext := filepath.Ext(fullpath)
	base := filepath.Base(fullpath)
	baseless := strings.TrimSuffix(base, ext)
	fullNewDir := filepath.Join(dir, baseless)
	createDir(fullNewDir)
	wg.Add(len(groups))
	for i, grp := range groups {
		stri := strconv.FormatInt(int64(i), 10)
		stringPath := baseless + "_" + stringTime() + "_" + stri + ext
		newPath := filepath.Join(fullNewDir, stringPath)
		go grp.initGroup(x, y, pixels, newPath)
	}
	wg.Wait()
}

func (grp group) initGroup(x int, y int, pixels allModifiedrgba, path string) {
	defer wg.Done()
	nx, ny, mod := grp.executeGroup(x, y, pixels)
	newimg := modify(nx, ny, mod)
	writeToFile(newimg, path)
}

func (grp group) executeGroup(x int, y int, pixels allModifiedrgba) (int, int, allModifiedrgba) {
	length := len(grp)
	pixelsAdd := pixels
	nx := x
	ny := y

	for x := 0; x < length; x++ {
		opt := strings.Split(grp[x], "=")
		switch opt[0] {
		case "d":
			//desaturate()
			option := opt[1]
			pixelsAdd = pixelsAdd.desaturate(option, desatFormula{0, 0, 0})
		case "df":
			//desaturateFormula()
			option := "custom"
			formula := strings.Split(opt[1], "_")
			// missing error !!!
			r, _ := strconv.ParseFloat(formula[0], 64)
			g, _ := strconv.ParseFloat(formula[1], 64)
			b, _ := strconv.ParseFloat(formula[2], 64)
			rgb := desatFormula{r, g, b}
			pixelsAdd = pixelsAdd.desaturate(option, rgb)
		case "b":
			//brightness()
			adjustment, _ := strconv.ParseFloat(opt[1], 64)
			pixelsAdd = pixelsAdd.brightness(adjustment)
		case "c":
			//contrast()
			adjustment, _ := strconv.ParseFloat(opt[1], 64)
			pixelsAdd = pixelsAdd.contrast(adjustment)
		case "h":
			//highlights()
			adjustment, _ := strconv.ParseFloat(opt[1], 64)
			pixelsAdd = pixelsAdd.highlights(adjustment)
		case "s":
			//shadows()
			adjustment, _ := strconv.ParseFloat(opt[1], 64)
			pixelsAdd = pixelsAdd.shadows(adjustment)
		case "bx":
			//border()
			borderOptions := strings.Split(opt[1], "_")
			borderWidth, _ := strconv.ParseInt(borderOptions[0], 10, 32)
			if len(borderOptions) == 1 {
				nx, ny, pixelsAdd = pixelsAdd.border(nx, ny, int(borderWidth), modifiedrgba{0, 0, 0, 0})
			} else if len(borderOptions) == 4 {
				r, _ := strconv.ParseUint(borderOptions[1], 10, 8)
				g, _ := strconv.ParseUint(borderOptions[2], 10, 8)
				b, _ := strconv.ParseUint(borderOptions[3], 10, 8)
				rgb := modifiedrgba{uint8(r), uint8(g), uint8(b), 255}
				nx, ny, pixelsAdd = pixelsAdd.border(nx, ny, int(borderWidth), rgb)
			}

		}
	}

	return nx, ny, pixelsAdd
}
