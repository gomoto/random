package main

import (
	"fmt"

	"github.com/gomoto/random/hex"
)

func main() {
	fraction := hex.Divide("AAAFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFF")
	fmt.Println(fraction)
}
