// Prints the boiling point of water. Let’s start getting into the
// habit of commenting what each of our files does going forward…
package main

import (
	"fmt"
)

// Because this is declared outside of a function, it is a PACKAGE-LEVEL
// DECLARATION! Not only can other members of this file access it, but so
// too can any other files in this entire package!
const boilingF = 212.0

func main() {
	// These variable declarations are accessible only to this function.
	// Pretty standard, really…
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Boiling point = %g°F or %g°C", f, c)
}
