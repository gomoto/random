package main

import (
	"fmt"
)

func main() {
	fraction := DivideHex("AAAFFFFFFFFFFFFF", "FFFFFFFFFFFFFFFF")
	fmt.Println(fraction)
}
