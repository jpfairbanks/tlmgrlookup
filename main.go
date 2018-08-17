package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

//go:generate ./assemble
func main() {
	arg := os.Args[1]
	fmt.Printf("Looking up latex package: %s\n", arg)
	filelist := strings.Split(Files, "\n")
	matches := 0
	for i, file := range filelist {
		if strings.Contains(file, arg) {
			matches += 1
			dir := path.Dir(file)
			pkg := path.Base(dir)
			fmt.Printf("Result: %d\t%s\t%s\n", i, file, pkg)
			fmt.Printf("\n\t sudo tlmgr install %s\n\n ", pkg)
		}
	}
	if matches == 0 {
		fmt.Printf("No Matches for %s\n", arg)
		os.Exit(1)
	}
}
