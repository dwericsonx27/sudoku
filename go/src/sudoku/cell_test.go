package sudoku

import (
	"fmt"
	"testing"
)

func TestCellCreation(t *testing.T) {
	var value = 8
	c, e := NewCell(value, 9)

	if e != nil {
		t.Errorf("ERROR creating cell object with value %d.", value)
	}

	expectedValue := c.GetValue()

	if expectedValue != value {
		t.Errorf("Expected value of %d, but it was %d instead.", value, expectedValue)
	}
	if c.NumPossibilities() != 0 {
		t.Errorf("ERROR: new cell contains %d possibilities when it should have 0.", c.NumPossibilities())
	}
	if c.GetMaxValue() != 9 {
		t.Errorf("Expected Max value not 9!")
	}
}

func TestCellCreationFailure(t *testing.T) {
	var value = 100
	var c CellInterface
	var e error
	c, e = NewCell(value, 9)

	if e == nil {
		t.Errorf("ERROR creating cell object with value %d should have caused an error.", value)
	}

	if c != nil {
		t.Errorf("ERROR: new cell object with invalid value should not be return.")
	}

	var maxValue = 8
	c, e = NewCell(value, maxValue)

	if e == nil {
		t.Errorf("Error not returned when creating cell object with invalid max value: %d.", maxValue)
	}

	if c != nil {
		t.Errorf("Error created with invalid max value: %d!", maxValue)
	}
}

func TestCellUnknownValuePossibilites(t *testing.T) {
	var c CellInterface
	var e error
	c, e = NewCell(-1, 9)
	if e == nil {
		for possibility := 1; possibility <= 9; possibility++ {
			if !(c.Contains(possibility)) {
				t.Errorf("ERROR: new cell should contain 1 as a possible value.")
			}
		}
		if c.NumPossibilities() != 9 {
			t.Errorf("ERROR: new cell contains %d possibilities when it should have 9.", c.NumPossibilities())
		}
	} else {
		t.Errorf("ERROR: failed to create cell.")
	}
}

func TestCellDetermined(t *testing.T) {
	c, e := NewCell(8, 9)

	if c != nil {
		if !c.Determined() {
			t.Errorf("Value for cell set, however reported as not set.")
		}
	} else {
		t.Errorf(e.Error())
	}
}

func TestCellValid(t *testing.T) {
	c, e := NewCell(8, 9)

	if c != nil {
		status, cellError := c.IsValid()
		if !status {
			t.Errorf(cellError.Error())
		}
	} else {
		t.Errorf(e.Error())
	}
}

func TestCandidates(t *testing.T) {
	c, e := NewCell(-1, 9)
	if c != nil {
		candidates := []int{2, 3, 5, 7}
		numCandidates := len(candidates)
		c.SetCandidates(candidates)
		s := c.GetCandidates()
		if s.Size() != numCandidates {
			msg := fmt.Sprintf("Number of candidates not %d.", numCandidates)
			t.Errorf(msg)
		} else {
			for i := 0; i > numCandidates; i++ {
				if !s.Contains(i) {
					msg := fmt.Sprintf("Expect candidate not found: %d", i)
					t.Errorf(msg)
				}
			}
		}
		if numCandidates > 1 {
			if c.GetValue() != -1 {
				t.Errorf("Cell value set when more than 1 candidate possible.")
			}
		}
	} else {
		t.Errorf(e.Error())
	}
}

func TestSingleCandidate(t *testing.T) {
	c, e := NewCell(-1, 9)
	if c != nil {
		expectedValue := 2
		candidates := []int{expectedValue}
		c.SetCandidates(candidates)
		actualValue := c.GetValue()
		if actualValue != expectedValue {
			msg := fmt.Sprintf("Only 1 candidate value set, expected value: %d, actual value %d.", expectedValue, actualValue)
			t.Errorf(msg)
		}
	} else {
		t.Errorf(e.Error())
	}
}
