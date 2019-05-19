package numbers

import (
	"fmt"
	"math/big"
)

// type alias
type Int = big.Int

type IntRange struct {
	Min *Int
	Max *Int
}

type IntRangeError struct {
	IntRange *IntRange
}

func (e *IntRangeError) Error() string {
	return fmt.Sprintf("Min (%s) cannot be greater than Max (%s)", e.IntRange.Min.String(), e.IntRange.Max.String())
}

type MapIntInputMinError struct {
	Value *Int
	Min   *Int
}

func (e *MapIntInputMinError) Error() string {
	return fmt.Sprintf("Value (%s) cannot be less than Min (%s)", e.Value.String(), e.Min.String())
}

type MapIntInputMaxError struct {
	Value *Int
	Max   *Int
}

func (e *MapIntInputMaxError) Error() string {
	return fmt.Sprintf("Value (%s) cannot be greater than Max (%s)", e.Value.String(), e.Max.String())
}

func MapInt(value Int, fromRange IntRange, toRange IntRange) (*Int, error) {
	if fromRange.Min.Cmp(fromRange.Max) == 1 {
		return nil, &IntRangeError{IntRange: &fromRange}
	}
	if toRange.Min.Cmp(toRange.Max) == 1 {
		return nil, &IntRangeError{IntRange: &toRange}
	}
	if value.Cmp(fromRange.Min) == -1 {
		return nil, &MapIntInputMinError{Value: &value, Min: fromRange.Min}
	}
	if value.Cmp(fromRange.Max) == 1 {
		return nil, &MapIntInputMaxError{Value: &value, Max: fromRange.Max}
	}
	// do math as floats

	outputValue := int64(0)
	return big.NewInt(outputValue), nil
}
