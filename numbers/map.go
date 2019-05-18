package numbers

import (
	"fmt"
)

type Uint64Range struct {
	Min uint64
	Max uint64
}

type Uint64RangeError struct {
	uint64Range Uint64Range
}

func (e *Uint64RangeError) Error() string {
	return fmt.Sprintf("min (%d) cannot be greater than max (%d)", e.uint64Range.Min, e.uint64Range.Max)
}

type MapUint64InputMinError struct {
	value uint64
	min   uint64
}

func (e *MapUint64InputMinError) Error() string {
	return fmt.Sprintf("value (%d) cannot be less than min (%d)", e.value, e.min)
}

type MapUint64InputMaxError struct {
	value uint64
	max   uint64
}

func (e *MapUint64InputMaxError) Error() string {
	return fmt.Sprintf("value (%d) cannot be less than min (%d)", e.value, e.max)
}

func MapUint64(value uint64, fromRange Uint64Range, toRange Uint64Range) (uint64, error) {
	if fromRange.Min > fromRange.Max {
		return 0, &Uint64RangeError{uint64Range: fromRange}
	}
	if toRange.Min > toRange.Max {
		return 0, &Uint64RangeError{uint64Range: toRange}
	}
	if value < fromRange.Min {
		return 0, &MapUint64InputMinError{value: value, min: fromRange.Min}
	}
	if fromRange.Max < value {
		return 0, &MapUint64InputMaxError{value: value, max: fromRange.Max}
	}
	// do math as floats

	outputValue := 0
	return uint64(outputValue), nil
}
