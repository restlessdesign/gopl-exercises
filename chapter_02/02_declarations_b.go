// Prints a few Celsius to Fahrenheit conversions.
//
// Why not say this package converts between degrees? It *does* (technically),
// but in a more proper setting, that logic should have its own package.
package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
