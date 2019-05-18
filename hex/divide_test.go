package hex_test

import (
	"testing"

	"github.com/gomoto/random/hex"
)

func TestDivide_NormalCase(t *testing.T) {
	hex.Divide("ABC", "FFF")
}

func TestDivide_NumeratorTooBig(t *testing.T) {
	hex.Divide("ABC", "FFF")
	// Divide("FFFFFFFFFFFFFFFFF", "FFF")
}
