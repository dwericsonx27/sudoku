package sudoku

import (
	"fmt"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	var val = 4
	if !(IsPerfectSquare(val)) {
		msg := fmt.Sprintf("%d reported not a perfect square!", val)
		t.Errorf(msg)
	}
}

func TestImperfectSquare(t *testing.T) {
	var val = 5
	if IsPerfectSquare(val) {
		msg := fmt.Sprintf("%d reported as a perfect square!", val)
		t.Errorf(msg)
	}
}

func checkRoot(t *testing.T, value int, expectedRoot int) {
	r, e := IntSquareRoot(value)
	if e != nil {
		if value >= 0 {
			errMsg := fmt.Sprintf("Integer square root of %d caused an unexpected error: %s", value, e.Error())
			t.Error(errMsg)
		}
	} else if expectedRoot != r {
		errMsg := fmt.Sprintf("Integer square root of %d expected to be %d, however received %d instead!", value, expectedRoot, r)
		t.Error(errMsg)
	}
}

func TestIntSquareRoot(t *testing.T) {
	checkRoot(t, 5, 2)
	checkRoot(t, 4, 2)
	checkRoot(t, 7, 2)
	checkRoot(t, -1, 0)
}
