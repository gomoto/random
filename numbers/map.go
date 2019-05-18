package numbers

import "fmt"

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

// Pass ranges by value. They are immutable.
func MapUint64(value uint64, fromRange Uint64Range, toRange Uint64Range) (uint64, error) {
	if fromRange.Min > fromRange.Max {
		return 0, &Uint64RangeError{uint64Range: fromRange}
	}
	if toRange.Min > toRange.Max {
		return 0, &Uint64RangeError{uint64Range: toRange}
	}
	return 0, nil
}
