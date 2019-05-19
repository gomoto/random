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
	toRangeMin := flag.Int64("min", 0, "Minimum integer (default 0)")
	toRangeMax := flag.Int64("max", 100, "Maximum integer (default 100)")
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
	intValue := new(big.Int)
	intMin := new(big.Int)
	intMax := new(big.Int)
	intValue.SetString(hexValue, 16)
	intMin.SetString(hexMin, 16)
	intMax.SetString(hexMax, 16)
	fromRange := numbers.IntRange{Min: intMin, Max: intMax}
	toRange := numbers.IntRange{Min: big.NewInt(*toRangeMin), Max: big.NewInt(*toRangeMax)}
	outputValue, err := numbers.MapInt(*intValue, fromRange, toRange)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err.Error())
	}
	fmt.Println(outputValue)
}
