package main

import (
	"fmt"
	"math/big"

	"github.com/gomoto/random/numbers"
)

func main() {
	inputValue := big.NewInt(3)
	fromRange := numbers.IntRange{Min: big.NewInt(2), Max: big.NewInt(4)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(19)}
	outputValue, err := numbers.MapInt(*inputValue, fromRange, toRange)
	fmt.Println(outputValue)
	fmt.Println(err)
}
