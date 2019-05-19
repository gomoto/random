package main

import (
	"fmt"
	"math/big"

	"github.com/gomoto/random/numbers"
)

func main() {
	inputValue := big.NewInt(17)
	fromRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	outputValue, err := numbers.MapInt(*inputValue, fromRange, toRange)
	fmt.Println(outputValue)
	fmt.Println(err)

}
