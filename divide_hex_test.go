package main_test

import (
	"testing"

	. "github.com/gomoto/random"
)

func TestDivideHex_NormalCase(t *testing.T) {
	DivideHex("ABC", "FFF")
}
