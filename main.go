package main

import (
	"fmt"

	"github.com/gomoto/random/numbers"
)

func main() {
	inputValue := uint64(17)
	fromRange := numbers.Uint64Range{Min: 10, Max: 20}
	toRange := numbers.Uint64Range{Min: 10, Max: 20}
	outputValue, err := numbers.MapUint64(inputValue, fromRange, toRange)
	fmt.Println(outputValue)
	fmt.Println(err)
}
