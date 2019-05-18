package hex_test

import (
	"testing"

	"github.com/gomoto/random/hex"
)

func TestDivideHex_NormalCase(t *testing.T) {
	hex.DivideHex("ABC", "FFF")
}

func TestDivideHex_NumeratorTooBig(t *testing.T) {
	hex.DivideHex("ABC", "FFF")
	// DivideHex("FFFFFFFFFFFFFFFFF", "FFF")
}
