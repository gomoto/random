package numbers_test

import (
	"math/big"
	"testing"

	"github.com/gomoto/random/numbers"
)

func TestMap_SuccessCase(t *testing.T) {
	inputValue := big.NewInt(3)
	fromRange := numbers.IntRange{Min: big.NewInt(2), Max: big.NewInt(4)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	outputValue, err := numbers.MapInt(*inputValue, fromRange, toRange)
	if err != nil {
		t.Error("error should be null")
	}
	expectedOutputValue := int64(15)
	actualOutputValue := outputValue.Int64()
	if actualOutputValue != expectedOutputValue {
		t.Errorf("output value is %d but was expected to be %d", actualOutputValue, expectedOutputValue)
	}
}

func TestMap_InvalidInputRange(t *testing.T) {
}

func TestMap_InvalidOutputRange(t *testing.T) {
}

func TestMap_ValueTooSmall(t *testing.T) {
}

func TestMap_ValueTooLarge(t *testing.T) {
}
