package sudoku

import (
	"bytes"
	"errors"
	"fmt"
	"set"
	"strings"
)

// Board is an array of Boxes to represent the Sudoku board.
//   <--- dimension in Boxes --->
//   ----------------------------
//   |        |        |        |
//   |  Box   |  Box   |  Box   |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Box   |  Box   |  Box   |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Box   |  Box   |  Box   |
//   |        |        |        |
//   ----------------------------
// The sudoku board shape is square.  The board consists of boxes arranged to make a square.
// A box consist of cells, arranged to make a square.  A cell holds a value, or a list
// of possible values.  The number of cells in a box, is the same as the number of boxes
// in a board.  The dimensions of a box in cells is also the same as the dimension of a board
// in boxes.  Subsequently the number of rows, columns or boxes in a board or the number of
// cells in a box must be a perfect square number.
type Board struct {
	boxes            map[int]BoxInterface
	dimensionInBoxes int
	maxValue         int
}

// NewBoard creates a Board object consisting of Boxes and Cells to represent a Sudoku board.
func NewBoard(dimensionSizeInBoxes int, boxConstructor func(sz int, CellConstructor func(value int, maxValue int) (CellInterface, error)) (BoxInterface, error), CellConstructor func(value int, maxValue int) (CellInterface, error)) (*Board, error) {
	if dimensionSizeInBoxes < 2 {
		msg := fmt.Sprintf("Board size must be 2 or greater, not: %d!", dimensionSizeInBoxes)
		return nil, errors.New(msg)
	}
	var err error
	var maxValue = dimensionSizeInBoxes * dimensionSizeInBoxes
	rtnval := &Board{make(map[int]BoxInterface), dimensionSizeInBoxes, maxValue}
	for i := 1; i <= maxValue; i++ {
		rtnval.boxes[i], err = boxConstructor(dimensionSizeInBoxes, CellConstructor)
		if err != nil {
			return nil, err
		}
	}
	return rtnval, nil
}

// NewBoardInitialize creates a new board based on known values.
func NewBoardInitialize(boardValues [][]int) (*Board, error) {
	numRows := len(boardValues)

	if IsPerfectSquare(numRows) {
		dimensionSizeInBoxes, _ := IntSquareRoot(numRows)
		b, e := NewBoard(dimensionSizeInBoxes, NewBox, NewCell)
		if e == nil {
			for rowIndex, rowElement := range boardValues {
				numColumns := len(rowElement)
				if numColumns != numRows {
					msg := fmt.Sprintf("Size of row(%d) is %d which does not match number of rows (%d)!",
						(rowIndex + 1), numColumns, numRows)
					return nil, errors.New(msg)
				}
				for colIndex, cellValue := range rowElement {
					b.SetValue(colIndex+1, rowIndex+1, int(cellValue))
				}
			}
			return b, nil
		}
		msg := fmt.Sprintf("Failed to create board!")
		return nil, errors.New(msg)
	}
	msg := fmt.Sprintf("Non square row count(%d)!", numRows)
	return nil, errors.New(msg)
}

func (b *Board) columnRowToBoxNum(column int, row int) (int, int, int) {
	var baseColumn = (column + 2) / b.dimensionInBoxes
	var offset = b.dimensionInBoxes * ((row - 1) / b.dimensionInBoxes)
	var boxNum = baseColumn + offset
	var boxColumn = 1 + ((column - 1) % b.dimensionInBoxes)
	var boxRow = 1 + ((row - 1) % b.dimensionInBoxes)
	return boxNum, boxColumn, boxRow
}

func (b *Board) getCell(column int, row int) (CellInterface, error) {
	boxNum, boxColumn, boxRow := b.columnRowToBoxNum(column, row)
	cell, e := b.boxes[boxNum].GetCell(boxColumn, boxRow)
	return cell, e
}

// SetValue sets the solved value at the location of a particular cell.
func (b *Board) SetValue(column int, row int, value int) error {
	cell, e := b.getCell(column, row)
	if e == nil {
		return cell.SetValue(value)
	}
	return e
}

// GetValue retreives the value set at the particular cell of the board.
func (b *Board) GetValue(column int, row int) (int, error) {
	cell, e := b.getCell(column, row)
	if e == nil {
		return cell.GetValue(), e
	}
	return -1, e //TBD: Need a better value than this.
}

// GetRepresentation returns an 2d array of integers representing the current cell values.
func (b *Board) GetRepresentation() ([][]int, error) {
	// Create the 2d array
	boardDimension := b.dimensionInBoxes * b.dimensionInBoxes
	board := make([][]int, boardDimension)
	for col := 0; col < boardDimension; col++ {
		board[col] = make([]int, boardDimension)
	}

	for col := 1; col <= b.maxValue; col++ {
		for row := 1; row <= b.maxValue; row++ {
			v, e := b.GetValue(col, row)
			if e == nil {
				board[row-1][col-1] = v
			} else {
				board[row-1][col-1] = -1
				msg := fmt.Sprintf("Failed to get cell at: (%d, %d)\n", col, row)
				return nil, errors.New(msg)
			}
		}
	}

	return board, nil
}

func (b *Board) findValuesColumn(column int, usedValues *set.IntSet) error {
	if 1 <= column && column <= b.maxValue {
		for row := 1; row <= b.maxValue; row++ {
			v, e := b.GetValue(column, row)
			if e == nil {
				if 1 <= v && v <= b.maxValue {
					usedValues.Add(v)
				}
			} else {
				msg := fmt.Sprintf("Board.findValuesColumn - cell not found at (%d, %d)\n", column, row)
				return errors.New(msg)
			}
		}
	}
	msg := fmt.Sprintf("Board.findValuesRow - invalid column(%d)\n", column)
	return errors.New(msg)
}

func (b *Board) findValuesRow(row int, usedValues *set.IntSet) error {
	if 1 <= row && row <= b.maxValue {
		for column := 1; column <= b.maxValue; column++ {
			v, e := b.GetValue(column, row)
			if e == nil {
				if 1 <= v && v <= b.maxValue {
					usedValues.Add(v)
				}
			} else {
				msg := fmt.Sprintf("Board.findValuesRow - cell not found at (%d, %d)\n", column, row)
				return errors.New(msg)
			}
		}
		return nil
	}
	msg := fmt.Sprintf("Board.findValuesRow - invalid row(%d)\n", row)
	return errors.New(msg)
}

func (b *Board) findValuesBox(column int, row int, usedValue *set.IntSet) error {
	if 1 <= column && column <= b.maxValue {
		if 1 <= row && row <= b.maxValue {
			boxNum, _, _ := b.columnRowToBoxNum(column, row)
			box := b.boxes[boxNum]
			for index := 1; index <= b.maxValue; index++ {
				cell, err := box.GetCellFromNum(index)
				if err != nil {
					return err
				}
				v := cell.GetValue()
				if 1 <= v && v <= b.maxValue {
					usedValue.Add(v)
				}
			}
			return nil
		}
	}
	msg := fmt.Sprintf("Board.findValuesBox - invalid location(%d, %d)\n", column, row)
	return errors.New(msg)
}

// FindHiddenSingle looks at the box, column and row that the specified cell is part of
// and attempts to determine the cell value.  If determination is not possible, then
// it eliminates possible values from the possiblities list for that cell.
func (b *Board) FindHiddenSingle(column int, row int) bool {
	cell, e := b.getCell(column, row)

	if e == nil {
		v := cell.GetValue()
		if 1 <= v && v <= cell.GetMaxValue() {
			return false // cell already solved.
		}

		usedValues := set.NewIntSet()
		b.findValuesColumn(column, usedValues)
		b.findValuesRow(row, usedValues)
		b.findValuesBox(column, row, usedValues)

		return cell.DiscardAndSetValue(usedValues)
	}
	return false
}

// FindNakedPairRow looks for a Naked Pair in a row, and eliminates these
// candidates from the other members of the row.
func (b *Board) FindNakedPairRow(column int, row int) bool {
	rtnval := false
	subjectCell, subjectError := b.getCell(column, row)

	matchingCellColumn := 0
	matchingCellCnt := 0
	if subjectError == nil {
		if subjectCell.NumPossibilities() == 2 {
			// Does a row contain a Naked Pair
			for columnIndex := 1; columnIndex <= b.maxValue; columnIndex++ {
				// Don't look at subject cell.
				if columnIndex != column {
					cell, e := b.getCell(columnIndex, row)
					if e == nil {
						if subjectCell.Equals(cell) {
							matchingCellColumn = columnIndex
							matchingCellCnt++
						}
					}
				}
			}
			// If so, eliminate values from other members of the row.
			if matchingCellCnt == 1 {
				for columnIndex := 1; columnIndex <= b.maxValue; columnIndex++ {
					// Don't look at subject cell, or the cell found to be the pair.
					if columnIndex != column && columnIndex != matchingCellColumn {
						cell, e := b.getCell(columnIndex, row)
						if e == nil {
							cell.DiscardAndSetValue(subjectCell.GetCandidates())
							rtnval = true
						}
					}
				}
			}
		}
	}
	return rtnval
}

// FindNakedPairColumn looks for a Naked Pair in a column, and eliminates these
// candidates from the other members of the column.
func (b *Board) FindNakedPairColumn(column int, row int) bool {
	rtnval := false
	subjectCell, subjectError := b.getCell(column, row)

	matchingCellRow := 0
	matchingCellCnt := 0
	if subjectError == nil {
		if subjectCell.NumPossibilities() == 2 {
			// Does a row contain a Naked Pair
			for rowIndex := 1; rowIndex <= b.maxValue; rowIndex++ {
				// Don't look at subject cell.
				if rowIndex != row {
					cell, e := b.getCell(column, rowIndex)
					if e == nil {
						if subjectCell.Equals(cell) {
							matchingCellRow = rowIndex
							matchingCellCnt++
						}
					}
				}
			}
			// If so, eliminate values from other members of the row.
			if matchingCellCnt == 1 {
				for rowIndex := 1; rowIndex <= b.maxValue; rowIndex++ {
					// Don't look at subject cell, or the cell found to be the pair.
					if rowIndex != row && rowIndex != matchingCellRow {
						cell, e := b.getCell(column, rowIndex)
						if e == nil {
							cell.DiscardAndSetValue(subjectCell.GetCandidates())
							rtnval = true
						}
					}
				}
			}
		}
	}
	return rtnval
}

// FindNakedPairBox looks for a Naked Pair in a box, and eliminates these
// candidates from the other members of the box.
func (b *Board) FindNakedPairBox(column int, row int) bool {
	boxNum, boxColumn, boxRow := b.columnRowToBoxNum(column, row)
	box := b.boxes[boxNum]
	return box.FindNakedPair(boxColumn, boxRow)
}

// FindNakedPair looks for the same pair in a row, column or box.
// If a naked pair is found then the pair values are eliminated from
// cells of the same structure the pair was found in.  For example if
// a naked pair is found in a row, then all other members of the row
// should have the values of the naked pair removed.
func (b *Board) FindNakedPair(column int, row int) bool {
	foundInRow := b.FindNakedPairRow(column, row)
	foundInColumn := b.FindNakedPairColumn(column, row)
	foundInBox := b.FindNakedPairBox(column, row)
	return foundInRow || foundInColumn || foundInBox
}

// SinglePassSolve steps through all cells of the Sudoku board and attemps to
// resolve the value for each cell.
func (b *Board) SinglePassSolve() bool {
	rtnval := false
	for col := 1; col <= b.maxValue; col++ {
		for row := 1; row <= b.maxValue; row++ {
			if b.FindHiddenSingle(col, row) {
				rtnval = true
			}

			if b.FindNakedPair(col, row) {
				rtnval = true
			}
		}
	}
	return rtnval
}

// AllCellsDetermined steps through every cell of the board to determine
// if all cells have had their value determined.
func (b *Board) AllCellsDetermined() bool {
	for col := 1; col <= b.maxValue; col++ {
		for row := 1; row <= b.maxValue; row++ {
			c, _ := b.getCell(col, row)
			if c != nil {
				if !c.Determined() {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func (b *Board) validColumn(column int) (bool, error) {
	catalog := make(map[int]int)
	for row := 1; row <= b.maxValue; row++ {
		value, _ := b.GetValue(column, row)
		if 1 <= value && value <= b.maxValue {
			_, exists := catalog[value]
			if exists {
				msg := fmt.Sprintf("%d already exists in column %d.", value, column)
				return false, errors.New(msg)
			}
			catalog[value] = 1
		}
	}
	return true, nil
}

func (b *Board) validRow(row int) (bool, error) {
	catalog := make(map[int]int)
	for column := 1; column <= b.maxValue; column++ {
		value, _ := b.GetValue(column, row)
		if 1 <= value && value <= b.maxValue {
			_, exists := catalog[value]
			if exists {
				msg := fmt.Sprintf("%d already exists in column %d.", value, column)
				return false, errors.New(msg)
			}
			catalog[value] = 1
		}
	}
	return true, nil
}

// IsValid steps through each row, column and box to make sure there is only
// 1 unique value other than the default value.
func (b *Board) IsValid() (bool, error) {
	for col := 1; col <= b.maxValue; col++ {
		ok, err := b.validColumn(col)
		if !ok {
			return false, err
		}
	}
	for row := 1; row <= b.maxValue; row++ {
		ok, err := b.validRow(row)
		if !ok {
			return false, err
		}
	}
	for boxNum := 1; boxNum <= b.maxValue; boxNum++ {
		ok, err := b.boxes[boxNum].IsValid()
		if !ok {
			return false, err
		}
	}
	return true, nil
}

// Solve keeps calling SingePassSolve till no more changes are made.  At this
// point the puzzle may NOT be solved because advanced puzzles may require searching
// a reduced problem space.
func (b *Board) Solve() bool {
	// Keep passing over the puzzle till no more changes are made.
	for b.SinglePassSolve() {
	}

	ok, _ := b.IsValid()
	return b.AllCellsDetermined() && ok
}

// Print sends an integer representation of the board to stdout.
func (b *Board) Print() {
	cnt := 4*b.maxValue + 1
	rowDivider := strings.Repeat("-", cnt)
	fmt.Printf("%s\n", rowDivider)
	for row := 1; row <= b.maxValue; row++ {
		for col := 1; col <= b.maxValue; col++ {
			v, _ := b.GetValue(col, row)
			fmt.Printf("|%2d ", v)
		}
		fmt.Printf("|\n")
		fmt.Printf("%s\n", rowDivider)
	}
}

// PrintPossibilities sends a representation to stdout of values of
// possible values for each cell.
func (b *Board) PrintPossibilities() {
	cnt := 4*b.maxValue + 1
	rowDivider := strings.Repeat("-", cnt)
	fmt.Printf("%s\n", rowDivider)
	for row := 1; row <= b.maxValue; row++ {
		for start := 1; start < b.maxValue; start += b.dimensionInBoxes {
			for col := 1; col <= b.maxValue; col++ {
				cell, _ := b.getCell(col, row)
				cellValue := cell.GetValue()
				if 1 <= cellValue && cellValue <= b.maxValue {
					stringValue := fmt.Sprintf("%d", cellValue)
					repeatedString := strings.Repeat(stringValue, b.dimensionInBoxes)
					fmt.Printf("|%s", repeatedString)
				} else {
					var buffer bytes.Buffer
					for v := start; v < start+b.dimensionInBoxes; v++ {
						if cell.Contains(v) {
							buffer.WriteString(fmt.Sprintf("%d", v))
						} else {
							buffer.WriteString(".")
						}
					}
					fmt.Printf("|%s", buffer.String())
				}
			}
			fmt.Printf("|\n")
		}
		fmt.Printf("%s\n", rowDivider)
	}
}

// SetCandidates sets the candidate values for a specified cell.
func (b *Board) SetCandidates(column int, row int, candidates []int) error {
	subjectCell, subjectError := b.getCell(column, row)
	if subjectCell != nil {
		subjectCell.SetCandidates(candidates)
	}
	return subjectError
}

// Equals determines if 2 boards are equivalent.
func (b Board) Equals(other *Board) bool {
	if b.maxValue == other.maxValue {
		for row := 1; row <= b.maxValue; row++ {
			for col := 1; col <= b.maxValue; col++ {
				c, e := b.getCell(col, row)
				if e == nil {
					c2, e2 := other.getCell(col, row)
					if e2 == nil {
						if !c.Equals(c2) {
							fmt.Printf("cells not equivavlent at %d, %d\n", col, row)
							return false
						}
					} else {
						fmt.Printf("error getting cell at %d, %d\n", col, row)
						return false
					}
				} else {
					fmt.Printf("error getting cell at %d, %d\n", col, row)
					return false
				}
			}
		}
		return true
	}
	fmt.Printf("boards not the same size\n")
	return false
}
