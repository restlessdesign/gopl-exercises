package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omit trailing newline")
var sep = flag.String("s", " ", "String separation character")

func main() {
	x := 1
	fmt.Println(x)

	// p “points” to the address of x (`&x`)
	p := &x

	// update the value at the address that p points to
	*p = 2
	fmt.Println(x)

	fmt.Println(x == *p)
	fmt.Println(&x == p)

	// Mismatched types!
	// fmt.Println(x == p)
	// fmt.Println(&x == &p)

	increment(p)
	fmt.Println(x)

	// The flag package uses pointers to convert command line arguments into
	// values our Go program can work with.
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// We can return an address from a function to create a closure of sorts
// *int meants that p’s type is a pointer to an int
func increment(p *int) int {
	// This will fail because you can’t perform math on a pointer type!
	// p += 1

	// …we need to extract the *value* of the pointer, p
	*p += 1

	return *p
}
