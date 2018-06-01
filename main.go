package main

import "fmt"

func main() {
	args := getArgs()
	all := parseArgs(args)
	fmt.Println(len(all))
	for _, x := range all {
		fmt.Println(x)
	}
}
