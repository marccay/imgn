package main

import (
	"strings"
	"sync"
)

var wg sync.WaitGroup

func execute(groups multipleGroups) {
	wg.Add(len(groups))
	for _, x := range groups {
		go executeGroup(x)
	}
	wg.Wait()
}

func executeGroup(grp group) {
	length := len(grp)

	for x := 0; x < length; x++ {
		opt := strings.Split(grp[x], "=")
		switch opt[0] {
		case "d":
			//desaturate()
		case "df":
			//desaturateFormula()
		case "b":
			//brightness()
		case "c":
			//contrast()
		}
	}
}
