package main

import (
	"fmt"
	"os"
)

// Main expects file name as the first argument (location 1).
// By default Args[0] contains the name of the executable.
func main() {
	fmt.Println("bundle: main{}")
	if len(os.Args) < 2 {
		fmt.Println("*** error: no input file ***")
		os.Exit(-1)
	}
	data := readAll(os.Args[1])
	unmarshal(data)
}