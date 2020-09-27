package sudoku

import (
	"errors"
	"fmt"
	"set"
)

// CellInterface describes the actions a cell can perform.
type CellInterface interface {
	SetValue(value int) error
	IsValid() (bool, error)
	GetValue() int
	GetMaxValue() int
	DiscardAndSetValue(usedValues *set.IntSet) bool
	Determined() bool
	Contains(possibility int) bool
	NumPossibilities() int
	GetCandidates() *set.IntSet
	Equals(cell CellInterface) bool
	SetCandidates(candidates []int)
}

// Cell is a structure containing the value, possiblities the value could be, and the
// maximum value.
type Cell struct {
	value         int
	possibilities *set.IntSet
	maxValue      int
}

// NewCell creates a cell object with a specified value and maximum value.
func NewCell(value int, maxValue int) (CellInterface, error) {
	c := Cell{}
	err := c.SetMaxValue(maxValue)
	if err != nil {
		return nil, err
	}
	err = c.SetValue(value)
	if err != nil {
		return nil, err
	}
	return CellInterface(&c), err
}

// GetValue returns the value specified in the Cell object.
func (c *Cell) GetValue() int {
	return c.value
}

// GetMaxValue returns the maximum value specified in the Cell object.
func (c *Cell) GetMaxValue() int {
	return c.maxValue
}

// SetMaxValue sets the maximum value in the Cell object.
func (c *Cell) SetMaxValue(maxValue int) error {
	var msg string
	if maxValue > 1 {
		if IsPerfectSquare(maxValue) {
			c.maxValue = maxValue
			return nil
		}
	}
	msg = fmt.Sprintf("Invalid cell max value: %d!", maxValue)
	return errors.New(msg)
}

// SetValue sets the value in the Cell object.
func (c *Cell) SetValue(value int) error {
	if (value >= 1) && (value <= c.maxValue) {
		c.value = value
		c.possibilities = set.NewIntSet()
	} else if value == -1 {
		c.value = value
		c.possibilities = set.NewIntSet()
		for i := 1; i <= 9; i++ {
			c.possibilities.Add(i)
		}
	} else {
		return errors.New("Invalid value for cell")
	}
	return nil
}

// Contains looks to see if the value in question still exists in the possiblities list.
func (c Cell) Contains(possibility int) bool {
	if (possibility >= 1) && (possibility <= c.maxValue) {
		return c.possibilities.Contains(possibility)
	}
	return false
}

// NumPossibilities return the number of possible values a cell could be.
func (c Cell) NumPossibilities() int {
	return c.possibilities.Size()
}

// IsNakedSingle is the situation where only 1 candidate value is remaining.
//
// NOTE THAT MEANS THE VALUE IS NOT SET YET, HOWEVER ONLY 1 CANDIDATE VALUE
// REMAINS.
func (c Cell) isNakedSingle() bool {
	return c.possibilities.Size() == 1
}

// DiscardAndSetValue eliminates values in the possiblities list based on an input
// list of values already in use.  If 1 value remains in the possiblities list, then
// that single value becomes the value of the cell.
func (c *Cell) DiscardAndSetValue(usedValues *set.IntSet) bool {
	if usedValues != nil {
		for _, v := range usedValues.GetAllMembers() {
			if c.possibilities.Contains(v) {
				c.possibilities.Remove(v)
			}
		}
	}
	if c.isNakedSingle() {
		fmt.Println("NakedSingle")
		val, e := c.possibilities.GetLastValue()
		if e == nil {
			c.SetValue(val)
			c.possibilities.Remove(val)
			return true
		}
	}
	return false
}

// Determined is true if the value for the cell has been set to a valid
// value.
func (c Cell) Determined() bool {
	if 1 <= c.value && c.value <= c.maxValue {
		return true
	}
	return false
}

// IsValid looks to see if the cell value has been determined, if so
// then the possiblities list should be empty, if not report false and an error.
func (c Cell) IsValid() (bool, error) {
	if 1 <= c.value && c.value <= c.maxValue {
		if c.possibilities.Size() == 0 {
			return true, nil
		}
		return false, errors.New("value set with remaining possibilities")
	} else if c.value == -1 {
		numPossibilities := c.possibilities.Size()
		if numPossibilities < 2 {
			return false, errors.New("value not set with 1 or less possibilities")
		} else if numPossibilities > c.maxValue {
			return false, errors.New("more possibilities than should exist")
		}
		return true, nil
	}
	msg := fmt.Sprintf("Cell value not valid: %d", c.value)
	return false, errors.New(msg)
}

// GetCandidates return the list of possibles candidate values for the subject cell.
func (c Cell) GetCandidates() *set.IntSet {
	return c.possibilities
}

// Equals determines if a pair of cells has equal candidates.
func (c Cell) Equals(subjectCell CellInterface) bool {
	c1 := c.GetCandidates()
	c2 := subjectCell.GetCandidates()
	return c1.Equals(c2)
}

// SetCandidates sets the specific set of candidates desired for a particular cell.
func (c *Cell) SetCandidates(candidates []int) {
	c.possibilities.Clear()
	for i := 0; i < len(candidates); i++ {
		c.possibilities.Add(candidates[i])
	}
	// if only 1 candidate, set value
	c.DiscardAndSetValue(nil)
}
