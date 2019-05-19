package numbers_test

import (
	"math/big"
	"testing"

	"github.com/gomoto/random/numbers"
)

func TestMap_SuccessCase(t *testing.T) {
	// Expect output value
	inputValue := int64(3)
	fromRange := numbers.IntRange{Min: big.NewInt(2), Max: big.NewInt(4)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	outputValue, err := numbers.MapInt(*big.NewInt(inputValue), fromRange, toRange)
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
	// Expect error IntRangeError when input range max is larger than input range min
	fromRangeMin := int64(6)
	fromRangeMax := int64(4)
	inputValue := int64(3)
	fromRange := numbers.IntRange{Min: big.NewInt(fromRangeMin), Max: big.NewInt(fromRangeMax)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	_, err := numbers.MapInt(*big.NewInt(inputValue), fromRange, toRange)
	// check underlying error type
	e, ok := err.(*numbers.IntRangeError)
	if !ok {
		t.Error("Expected error to have type IntRangeError")
	} else {
		if e.IntRange.Min.Int64() != fromRangeMin {
			t.Errorf("Expected Min to be %d", fromRangeMin)
		}
		if e.IntRange.Max.Int64() != fromRangeMax {
			t.Errorf("Expected Max to be %d", fromRangeMax)
		}
	}
}

func TestMap_InvalidOutputRange(t *testing.T) {
	// Expect error IntRangeError when output range max is larger than output range min
	toRangeMin := int64(20)
	toRangeMax := int64(10)
	inputValue := int64(3)
	fromRange := numbers.IntRange{Min: big.NewInt(2), Max: big.NewInt(4)}
	toRange := numbers.IntRange{Min: big.NewInt(toRangeMin), Max: big.NewInt(toRangeMax)}
	_, err := numbers.MapInt(*big.NewInt(inputValue), fromRange, toRange)
	// check underlying error type
	e, ok := err.(*numbers.IntRangeError)
	if !ok {
		t.Error("Expected error to have type IntRangeError")
	} else {
		if e.IntRange.Min.Int64() != toRangeMin {
			t.Errorf("Expected Min to be %d", toRangeMin)
		}
		if e.IntRange.Max.Int64() != toRangeMax {
			t.Errorf("Expected Max to be %d", toRangeMax)
		}
	}
}

func TestMap_ValueTooSmall(t *testing.T) {
	// Expect error MapIntInputMinError when input value is smaller than input range min
	inputValue := int64(1)
	fromRangeMin := int64(2)
	fromRange := numbers.IntRange{Min: big.NewInt(fromRangeMin), Max: big.NewInt(4)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	_, err := numbers.MapInt(*big.NewInt(inputValue), fromRange, toRange)
	// check underlying error type
	e, ok := err.(*numbers.MapIntInputMinError)
	if !ok {
		t.Error("Expected error to have type MapIntInputMinError")
	} else {
		if e.Min.Int64() != fromRangeMin {
			t.Errorf("Expected Min to be %d", fromRangeMin)
		}
		if e.Value.Int64() != inputValue {
			t.Errorf("Expected Value to be %d", inputValue)
		}
	}
}

func TestMap_ValueTooLarge(t *testing.T) {
	// Expect error MapIntInputMaxError when input value is larger than input range max
	inputValue := int64(5)
	fromRangeMax := int64(4)
	fromRange := numbers.IntRange{Min: big.NewInt(2), Max: big.NewInt(fromRangeMax)}
	toRange := numbers.IntRange{Min: big.NewInt(10), Max: big.NewInt(20)}
	_, err := numbers.MapInt(*big.NewInt(inputValue), fromRange, toRange)
	// check underlying error type
	e, ok := err.(*numbers.MapIntInputMaxError)
	if !ok {
		t.Error("Expected error to have type MapIntInputMaxError")
	} else {
		if e.Max.Int64() != fromRangeMax {
			t.Errorf("Expected Max to be %d", fromRangeMax)
		}
		if e.Value.Int64() != inputValue {
			t.Errorf("Expected Value to be %d", inputValue)
		}
	}
}
