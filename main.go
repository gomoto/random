package main

import (
	"fmt"

	"github.com/gomoto/random/hex"
)

func main() {
	fraction := hex.DivideHex("AAAFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFF")
	fmt.Println(fraction)
}
