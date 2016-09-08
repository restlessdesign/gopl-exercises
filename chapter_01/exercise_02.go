package main

import (
	"fmt"
	"os"
	"testing"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
}
