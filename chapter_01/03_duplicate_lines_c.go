package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Ignoring potential errors for input.Err() for now
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// ioutil.ReadFile only takes named files, so we can’t use stdin anymore
		// It also doesn’t require us to open a file anymore, so we don’t need
		// to make a call to os.Close(f)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s %v\n", file, err)
			continue
		}

		for _, line := range strings.Split(data, "\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
