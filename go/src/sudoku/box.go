package sudoku

import (
	"errors"
	"fmt"
)

// BoxInterface describes the interface to interact with Box object types.
type BoxInterface interface {
	GetCell(column int, row int) (CellInterface, error)
	GetCellFromNum(cellNum int) (CellInterface, error)
	IsValid() (bool, error)
	SetValue(column int, row int, value int) error
	FindNakedPair(boxColumn int, boxRow int) bool
}

// Box holds cells and meta data associated with its dimensionality.
//   <--- dimension in Cells --->
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------
type Box struct {
	cells            map[int]CellInterface
	dimensionInCells int
	maxValue         int
}

// NewBox creates a box object given a cell constructor.
func NewBox(dimensionInCells int, CellConstructor func(value int, maxValue int) (CellInterface, error)) (BoxInterface, error) {
	if dimensionInCells < 2 {
		msg := fmt.Sprintf("Box size must be 2 or greater, not: %d!", dimensionInCells)
		return nil, errors.New(msg)
	}
	var err error
	maxValue := dimensionInCells * dimensionInCells
	rtnval := &Box{make(map[int]CellInterface), dimensionInCells, maxValue}
	for i := 1; i <= maxValue; i++ {
		rtnval.cells[i], err = CellConstructor(-1, maxValue)
		if err != nil {
			return nil, err
		}
	}
	return BoxInterface(rtnval), nil
}

func (b *Box) validLocation(column int, row int) bool {
	if column >= 1 && column <= b.dimensionInCells {
		if row >= 1 && row <= b.dimensionInCells {
			return true
		}
	}
	return false
}

func (b *Box) colRowToCellNum(column int, row int) int {
	return column + (row-1)*3
}

// GetCell returns a reference to the desired cell within a Box.
func (b *Box) GetCell(column int, row int) (CellInterface, error) {
	if !b.validLocation(column, row) {
		return nil, errors.New("invalid column/row specified for GetCell()")
	}
	cellNum := b.colRowToCellNum(column, row)
	return b.cells[cellNum], nil
}

// SetValue sets the values of a cell in a Box.
func (b *Box) SetValue(column int, row int, value int) error {
	cell, e := b.GetCell(column, row)
	if e == nil {
		return cell.SetValue(value)
	}
	return e
}

// IsValid steps through the cells of the box making sure the values are
// unique or the default unset value.
func (b *Box) IsValid() (bool, error) {
	catalog := make(map[int]int)
	for cellNum := 1; cellNum <= b.maxValue; cellNum++ {
		c := b.cells[cellNum]
		cellOK, err := c.IsValid()
		if err != nil {
			return false, errors.New(err.Error())
		} else if cellOK {
			v := c.GetValue()
			if v >= 1 {
				_, exists := catalog[v]
				if exists {
					msg := fmt.Sprintf("%d exists more than once in box", v)
					return false, errors.New(msg)
				}
				catalog[v] = 1
			}
		} else {
			msg := fmt.Sprintf("Cell Number %d invalid in box.", cellNum)
			return false, errors.New(msg)
		}
	}
	return true, nil
}

// GetCellFromNum returns a reference to a cell based on its location in a 1d array.
func (b *Box) GetCellFromNum(cellNum int) (CellInterface, error) {
	if cellNum >= 1 && cellNum <= b.maxValue {
		return b.cells[cellNum], nil
	}
	msg := fmt.Sprintf("Invalid cell number : %d", cellNum)
	return nil, errors.New(msg)
}

// FindNakedPair looks for naked pairs in a box, and eliminates these values
// from each of the other cells of the box.
func (b *Box) FindNakedPair(boxColumn int, boxRow int) bool {
	rtnval := false
	subjectCellNum := b.colRowToCellNum(boxColumn, boxRow)
	subjectCell, subjectError := b.GetCellFromNum(subjectCellNum)
	if subjectError == nil {
		// Look to see if there are any naked pairs.
		matchCellNum := 0
		matchCellCnt := 0
		for i := 1; i <= b.maxValue; i++ {
			// Don't match the subject cell.
			if i != subjectCellNum {
				c, e := b.GetCellFromNum(i)
				if e == nil {
					if subjectCell.Equals(c) {
						matchCellNum = i
						matchCellCnt++
					}
				}
			}
		}

		if matchCellCnt == 1 {
			// Remove the pair of candidate values from each of the other cells of the box.
			for i := 1; i <= b.maxValue; i++ {
				if i != subjectCellNum && i != matchCellNum {
					c, e := b.GetCellFromNum(i)
					if e == nil {
						if c.DiscardAndSetValue(subjectCell.GetCandidates()) {
							rtnval = true
						}
					}
				}
			}
		}
	}
	return rtnval
}
