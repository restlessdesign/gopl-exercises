package main

import (
	"bufio"
	"fmt"
	"os"
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
	files := os.Args[1:]

	if len(files) == 0 {
		// No files passed; read from stdin
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s %v\n", file, err)
				continue
			}

			countLines(f, counts)

			// This should probably be called immediately using `defer`, but
			// the book hasn’t covered that topic yet :)
			f.Close()
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
