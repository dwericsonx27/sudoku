package sudoku

import (
	"errors"
	"math"
)

// IntSquareRoot return an integer of a square root that when square does not
// exceed the original value. Considered using Newton's method instead, this
// implementation was on the order of 6 times faster.
func IntSquareRoot(value int) (int, error) {
	if value >= 0 {
		return int(math.Floor(math.Sqrt(float64(value)))), nil
	}
	return 0, errors.New("integer square root of a negative number nonsensical")
}

// IsPerfectSquare determines if the value is a perfect square.
func IsPerfectSquare(value int) bool {
	intRoot, err := IntSquareRoot(value)
	if err != nil {
		return false
	}
	return (intRoot * intRoot) == value
}
