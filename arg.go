package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type group map[int]string
type multipleGroups []group

// getArgs gets raw info from given arguments
func getArgs() (qnty int64, path string, flags map[int]string) {
	args := os.Args[3:]
	mapFlags := make(map[int]string)
	quantity := os.Args[1]
	qnty, _ = strconv.ParseInt(quantity, 10, 64)
	for x := 0; x < int(qnty); x++ {
		stringX := strconv.FormatInt(int64(x+1), 10)
		for i, arg := range args {
			if arg == ("-" + stringX) {
				// Error
				// if option is cut off in ending
				if i+1 >= len(args) {
					fmt.Printf("not enough options described in output %d\n", x+1)
					os.Exit(2)
				}
				// Error
				// if option (i+1) is next flag
				if args[i+1][0] == 45 {
					fmt.Printf("not enough options descripted output %d\n", x+1)
					os.Exit(2)
				}
				mapFlags[x] = args[i+1]
			}
			if i == len(args) && arg != ("-"+stringX) {
				fmt.Printf("missing options for %d\n", x)
				os.Exit(2)
			}
		}
	}
	// Error
	// quantity doesn't match given input
	if int(qnty) != len(mapFlags) {
		fmt.Println("given quantity does not match given options")
		os.Exit(2)
	}

	path = filepath.Dir(os.Args[2])

	return qnty, path, mapFlags
}

func parseArgs(unpackedData map[int]string) multipleGroups {
	all := make(multipleGroups, len(unpackedData))
	for i, data := range unpackedData {
		g := strings.Split(data, ",")
		grp := make(group, len(g))
		for x, gr := range g {
			grp[x] = gr
		}
		all[i] = grp
	}
	return all
}
