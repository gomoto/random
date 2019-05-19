package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/gomoto/random/nist"
	"github.com/gomoto/random/numbers"
)

// Print random number between specified min and max
func main() {
	// Get target range from CLI args
	toRangeMinRaw := flag.Int64("min", 0, "Minimum integer (default 0)")
	toRangeMaxRaw := flag.Int64("max", 100, "Maximum integer (default 100)")
	flag.Parse()

	// Get latest random value
	hexValue, err := nist.GetLatestValue()
	if err != nil {
		fmt.Println("Error 1")
		fmt.Println(err.Error())
	}
	hexMin := nist.GetMinValue()
	hexMax := nist.GetMaxValue()

	// Convert random value to a value in the target range
	fromValue := new(big.Int)
	fromRangeMin := new(big.Int)
	fromRangeMax := new(big.Int)
	fromValue.SetString(hexValue, 16)
	fromRangeMin.SetString(hexMin, 16)
	fromRangeMax.SetString(hexMax, 16)
	toRangeMin := big.NewInt(*toRangeMinRaw)
	toRangeMax := big.NewInt(*toRangeMaxRaw)
	// Add one to each maximum before mapping
	fromRange := numbers.IntRange{Min: fromRangeMin, Max: fromRangeMax}
	toRange := numbers.IntRange{Min: toRangeMin, Max: toRangeMax}
	outputValue, err := numbers.MapInt(*fromValue, fromRange, toRange)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err.Error())
	}
	fmt.Println(outputValue)
}
