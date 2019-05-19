package main

import (
	"fmt"
	"math/big"

	"github.com/gomoto/random/numbers"
)

func main() {
	hexValue := "C60BF83BDCBAD2C73F76B5AAC4604603EB325D690969545A28EAD1880A159C8EBEE3E1647957CC9F7AB16F4E7457B40B95F86497A9AEE59BFAFCEFD660927235"
	hexMin := "0"
	hexMax := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	intValue := new(big.Int)
	intMin := new(big.Int)
	intMax := new(big.Int)
	intValue.SetString(hexValue, 16)
	intMin.SetString(hexMin, 16)
	intMax.SetString(hexMax, 16)
	fmt.Println(intValue)
	fmt.Println(intMin)
	fmt.Println(intMax)
	fromRange := numbers.IntRange{Min: intMin, Max: intMax}
	toRange := numbers.IntRange{Min: big.NewInt(1), Max: big.NewInt(49)}
	outputValue, _ := numbers.MapInt(*intValue, fromRange, toRange)
	fmt.Println(outputValue)
}
