package main

import (
	"fmt"
	"strconv"
)

func main() {
	fraction := divideHex("AAFFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFF")
	fmt.Println(fraction)
}

// Divide two base-16 numbers.
func divideHex(hexNumerator string, hexDenominator string) float64 {
	// If either hex value is larger than the maximum allowed
	// ("FFFFFFFFFFFFFFFF"), ParseUint will panic.
	intDenominator, err := strconv.ParseUint(hexDenominator, 16, 64)
	if err != nil {
		panic(err)
	}
	intNumerator, err := strconv.ParseUint(hexNumerator, 16, 64)
	if err != nil {
		panic(err)
	}
	fraction := float64(intNumerator) / float64(intDenominator)
	return fraction
}
