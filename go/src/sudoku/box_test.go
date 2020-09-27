package sudoku

import (
	"testing"
)

func TestBoxCreation(t *testing.T) {
	b, e := NewBox(3, NewCell)

	if e != nil {
		t.Errorf("ERROR creating box object")
		t.Errorf(e.Error())
	}

	if b == nil {
		t.Errorf("ERROR box object not returned.")
	}
}

func TestInvalidBoxCreation(t *testing.T) {
	_, e := NewBox(1, NewCell)
	if e == nil {
		t.Errorf("Invalid box size did not return an error.")
	}
}

func TestGetCellFromBox(t *testing.T) {
	var b BoxInterface
	var e error
	b, e = NewBox(3, NewCell)

	var c CellInterface
	c, e = b.GetCell(1, 1)

	if e != nil {
		t.Errorf("ERROR fetching cell at 1, 1 returned an error.")
		t.Errorf(e.Error())
	}

	if c == nil {
		t.Errorf("ERROR cell at 1, 1 not found.")
	}
	c.SetValue(6)

	var c2 CellInterface
	c2, e = b.GetCell(1, 1)

	if e == nil {
		if c2.GetValue() != 6 {
			t.Errorf("Value expected for cell 1, 1, not 6")
		}
	}

	c, e = b.GetCell(1, 4)
	if e == nil {
		t.Errorf("No error generated for invalid cell location.")
		t.Errorf(e.Error())
	}

	if c != nil {
		t.Errorf("Cell return for invalid coordinate.")
	}
}

func TestValidityOfBox(t *testing.T) {
	var size = 3
	box, eBox := NewBox(size, NewCell)

	if box != nil {
		for c := 1; c <= size; c++ {
			for r := 1; r <= size; r++ {
				value := c + 3*(r-1)
				eValue := box.SetValue(r, c, value)
				if eValue != nil {
					t.Errorf("Setting a cell value from box failed!")
					t.Errorf(eValue.Error())
				}
			}
		}
		ok, _ := box.IsValid()
		if !ok {
			t.Errorf("Valid box designated as invalid!")
		}

		if box.SetValue(2, 2, 2) == nil {
			ok, _ = box.IsValid()
			if ok {
				t.Errorf("Invalid box designated as valid!")
			}
		} else {
			t.Errorf("Failed to set 2 at (2, 2).")
		}
	} else {
		t.Errorf(eBox.Error())
	}
}
