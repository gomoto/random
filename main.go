package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/gomoto/random/numbers"
)

// Print random number between specified min and max
func main() {
	toRangeMin := flag.Int64("min", 0, "Minimum integer (default 0)")
	toRangeMax := flag.Int64("max", 0, "Maximum integer (default 0)")
	flag.Parse()
	hexValue := "C60BF83BDCBAD2C73F76B5AAC4604603EB325D690969545A28EAD1880A159C8EBEE3E1647957CC9F7AB16F4E7457B40B95F86497A9AEE59BFAFCEFD660927235"
	hexMin := "0"
	hexMax := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	intValue := new(big.Int)
	intMin := new(big.Int)
	intMax := new(big.Int)
	intValue.SetString(hexValue, 16)
	intMin.SetString(hexMin, 16)
	intMax.SetString(hexMax, 16)
	fromRange := numbers.IntRange{Min: intMin, Max: intMax}
	toRange := numbers.IntRange{Min: big.NewInt(*toRangeMin), Max: big.NewInt(*toRangeMax)}
	outputValue, _ := numbers.MapInt(*intValue, fromRange, toRange)
	fmt.Println(outputValue)
}
