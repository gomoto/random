package main

import (
	"fmt"
)

func main() {
	fraction := divideHex("AAAFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFF")
	fmt.Println(fraction)
}
