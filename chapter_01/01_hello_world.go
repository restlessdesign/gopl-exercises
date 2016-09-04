package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// The name of the file is technically the first argument passed, so skip
	// over os.Args[0]!
	//
	// $ go run 01_echo1.go foo bar baz
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
