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

// Map input integer to output integer, given input and output integer ranges.
func MapInt(value Int, fromRange IntRange, toRange IntRange) (*Int, error) {
	// Solve for Value2:
	// (Max2 - Min2)/(Max1 - Min1) = (Value2 - Min2)/(Value1 - Min1)
	// Value2 = Min2 + (Value1 - Min1)(Max2 - Min2)/(Max1 - Min1)
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
	// (Value1 - Min1)
	numeratorA := new(Int)
	numeratorA.Sub(&value, fromRange.Min)
	// (Max2 - Min2)
	numeratorB := new(Int)
	numeratorB.Sub(toRange.Max, toRange.Min)
	numerator := new(Int)
	numerator.Mul(numeratorA, numeratorB)
	// (Max1 - Min1)
	denominator := new(Int)
	denominator.Sub(fromRange.Max, fromRange.Min)
	ratProduct := new(big.Rat)
	ratProduct.SetFrac(numerator, denominator)
	floatProduct := new(big.Float)
	floatProduct.SetRat(ratProduct)
	intProduct, _ := floatProduct.Int(nil)
	outputValue := new(Int)
	// Add Min2
	outputValue.Add(intProduct, toRange.Min)
	return outputValue, nil
}
