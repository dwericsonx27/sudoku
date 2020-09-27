package sudoku

import (
	"errors"
	"fmt"
	"testing"
)

var solvableBoard1 = [][]int{
	{-1, -1, -1, 2, 6, -1, 7, -1, 1},
	{6, 8, -1, -1, 7, -1, -1, 9, -1},
	{1, 9, -1, -1, -1, 4, 5, -1, -1},
	{8, 2, -1, 1, -1, -1, -1, 4, -1},
	{-1, -1, 4, 6, -1, 2, 9, -1, -1},
	{-1, 5, -1, -1, -1, 3, -1, 2, 8},
	{-1, -1, 9, 3, -1, -1, -1, 7, 4},
	{-1, 4, -1, -1, 5, -1, -1, 3, 6},
	{7, -1, 3, -1, 1, 8, -1, -1, -1},
}

var solutionBoard1 = [][]int{
	{4, 3, 5, 2, 6, 9, 7, 8, 1},
	{6, 8, 2, 5, 7, 1, 4, 9, 3},
	{1, 9, 7, 8, 3, 4, 5, 6, 2},
	{8, 2, 6, 1, 9, 5, 3, 4, 7},
	{3, 7, 4, 6, 8, 2, 9, 1, 5},
	{9, 5, 1, 7, 4, 3, 6, 2, 8},
	{5, 1, 9, 3, 2, 6, 8, 7, 4},
	{2, 4, 8, 9, 5, 7, 1, 3, 6},
	{7, 6, 3, 4, 1, 8, 2, 5, 9},
}

func FakeNewBox(dimensionInCells int) (*Box, error) {
	return nil, errors.New("intentional error condition created")
}

func compare2dArrays(a [][]int, b [][]int) bool {
	var size = len(a)
	if size == len(b) {
		for c := 0; c < size; c++ {
			if len(a[c]) == len(b[c]) {
				for r := 0; r < size; r++ {
					if a[c][r] != b[c][r] {
						return false // elements do not match
					}
				}
			} else {
				return false // column sizes do not match
			}
		}
		return true // column and row sizes match along with element matching.
	}
	return false // overall sizes different
}

func TestBoardCreation(t *testing.T) {
	b, e := NewBoard(3, NewBox, NewCell)

	if e != nil {
		t.Errorf("ERROR creating board object")
		t.Errorf(e.Error())
	}

	if b == nil {
		t.Errorf("ERROR board object not returned.")
	}
}

func TestBoardCreationInvalidSize(t *testing.T) {
	b, e := NewBoard(1, NewBox, NewCell)

	if e == nil {
		t.Errorf("ERROR: not error detected for invalid size.")
	}

	if b != nil {
		t.Errorf("ERROR: board object returned for an invalid size.")
	}
}

func TestBoardInitializiation(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if e != nil {
		t.Error("Board not created!")
		t.Error(e)
	} else {
		board, e := b.GetRepresentation()
		if e != nil {
			t.Error("Failed to get board representation!!")
		}

		if !compare2dArrays(solvableBoard1, board) {
			t.Error("Input board does not match extracted board.")
		}
	}
}

func TestSetGetValue(t *testing.T) {
	b, e := NewBoard(3, NewBox, NewCell)

	if e != nil {
		t.Errorf("ERROR creating board object")
		t.Errorf(e.Error())
	} else {
		var testVal = 7
		b.SetValue(7, 3, testVal)
		v, _ := b.GetValue(7, 3)
		if testVal != v {
			msg := fmt.Sprintf("Set value %d, does not match received value %d\n", testVal, v)
			t.Errorf(msg)
		}
	}
}

func TestFindHiddenSingle(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if e == nil {
		if !b.FindHiddenSingle(3, 1) {
			t.Errorf("Failed to solve cell at 3, 1.")
		}
		v, _ := b.GetValue(3, 1)
		if v != 5 {
			msg := fmt.Sprintf("Expected solved value of 5 at 3, 1, instead %d!", v)
			t.Errorf(msg)
		}
	} else {
		t.Errorf("Failed to create board!")
	}
}

func TestSinglePassSolve(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if e == nil {
		if !b.SinglePassSolve() {
			t.Errorf("Fail to solve a single cell in board.")
		}
	} else {
		t.Errorf("Failed to create board!")
	}
}

func TestAllCellsDetermined(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if b != nil {
		if b.AllCellsDetermined() {
			t.Error("Incomplete board characterized as complete.")
		}
	} else {
		t.Errorf(e.Error())
	}
	b, e = NewBoardInitialize(solutionBoard1)
	if b != nil {
		if !b.AllCellsDetermined() {
			t.Error("Solved board characterized as incomplete.")
		}
	} else {
		t.Errorf(e.Error())
	}
}

func TestValid(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if e == nil {
		ok, _ := b.IsValid()
		if !ok {
			t.Errorf("Valid board board declared invalid.")
		}
	} else {
		t.Errorf("Failed to create board!")
	}
}

func TestSolve(t *testing.T) {
	b, e := NewBoardInitialize(solvableBoard1)
	if e == nil {
		if !b.Solve() {
			t.Errorf("Fail to solve a single cell in board.")
		} else {
			board, eBoard := b.GetRepresentation()
			if board != nil {
				if !compare2dArrays(solutionBoard1, board) {
					t.Error("Computed solution does not match solution.")
				}
			} else {
				t.Errorf(eBoard.Error())
			}
		}
	} else {
		t.Errorf("Failed to create board!")
	}
}

func TestNakedPair(t *testing.T) {
	dimensionSizeInBoxes := 3
	b, e := NewBoard(dimensionSizeInBoxes, NewBox, NewCell)
	if b != nil {
		b.SetCandidates(1, 1, []int{3, 5, 8})
		b.SetValue(2, 1, 7)
		b.SetValue(3, 1, 9)
		b.SetCandidates(4, 1, []int{2, 3, 4, 8})
		b.SetCandidates(5, 1, []int{2, 3, 4, 5})
		b.SetCandidates(6, 1, []int{2, 8})
		b.SetValue(7, 1, 6)
		b.SetCandidates(8, 1, []int{2, 4, 5, 8})
		b.SetValue(9, 1, 1)

		b.SetCandidates(1, 2, []int{3, 5, 8})
		b.SetCandidates(2, 2, []int{5, 8})
		b.SetValue(3, 2, 1)
		b.SetCandidates(4, 2, []int{2, 3, 4, 8})
		b.SetValue(5, 2, 6)
		b.SetValue(6, 2, 7)
		b.SetCandidates(7, 2, []int{2, 4, 9})
		b.SetCandidates(8, 2, []int{2, 4, 5, 8})
		b.SetCandidates(9, 2, []int{2, 4, 5, 8, 9})

		b.SetValue(1, 3, 6)
		b.SetValue(2, 3, 2)
		b.SetValue(3, 3, 4)
		b.SetValue(4, 3, 9)
		b.SetCandidates(5, 3, []int{1, 5})
		b.SetCandidates(6, 3, []int{1, 8})
		b.SetValue(7, 3, 3)
		b.SetValue(8, 3, 7)
		b.SetCandidates(9, 3, []int{5, 8})

		b.SetCandidates(1, 4, []int{1, 5, 8})
		b.SetCandidates(2, 4, []int{1, 5, 8})
		b.SetValue(3, 4, 6)
		b.SetCandidates(4, 4, []int{2, 3, 7, 8})
		b.SetCandidates(5, 4, []int{2, 3, 7, 9})
		b.SetValue(6, 4, 4)
		b.SetCandidates(7, 4, []int{1, 2, 7, 9})
		b.SetCandidates(8, 4, []int{1, 2, 3, 5, 8})
		b.SetCandidates(9, 4, []int{2, 5, 7, 8, 9})

		b.SetCandidates(1, 5, []int{4, 8})
		b.SetCandidates(2, 5, []int{4, 8})
		b.SetCandidates(3, 5, []int{2, 7})
		b.SetValue(4, 5, 1)
		b.SetCandidates(5, 5, []int{2, 3, 7, 9})
		b.SetValue(6, 5, 5)
		b.SetCandidates(7, 5, []int{2, 4, 7, 9})
		b.SetValue(8, 5, 6)
		b.SetCandidates(9, 5, []int{2, 4, 7, 8, 9})

		b.SetValue(1, 6, 9)
		b.SetValue(2, 6, 3)
		b.SetCandidates(3, 6, []int{2, 7})
		b.SetCandidates(4, 6, []int{2, 6, 7, 8})
		b.SetCandidates(5, 6, []int{2, 7})
		b.SetCandidates(6, 6, []int{2, 6, 8})
		b.SetCandidates(7, 6, []int{1, 2, 4, 7})
		b.SetCandidates(8, 6, []int{1, 2, 4, 5, 8})
		b.SetCandidates(9, 6, []int{2, 4, 5, 7, 8})

		b.SetCandidates(1, 7, []int{1, 2, 4})
		b.SetValue(2, 7, 6)
		b.SetValue(3, 7, 5)
		b.SetCandidates(4, 7, []int{2, 4, 7})
		b.SetCandidates(5, 7, []int{1, 2, 4, 7})
		b.SetCandidates(6, 7, []int{1, 2})
		b.SetValue(7, 7, 8)
		b.SetValue(8, 7, 9)
		b.SetValue(9, 7, 3)

		b.SetValue(1, 8, 7)
		b.SetValue(2, 8, 9)
		b.SetValue(3, 8, 8)
		b.SetCandidates(4, 8, []int{2, 4, 6})
		b.SetCandidates(5, 8, []int{1, 2, 4})
		b.SetValue(6, 8, 3)
		b.SetValue(7, 8, 5)
		b.SetCandidates(8, 8, []int{1, 2, 4})
		b.SetCandidates(9, 8, []int{2, 4, 6})

		b.SetCandidates(1, 9, []int{1, 2, 4})
		b.SetCandidates(2, 9, []int{1, 4})
		b.SetValue(3, 9, 3)
		b.SetValue(4, 9, 5)
		b.SetValue(5, 9, 8)
		b.SetValue(6, 9, 9)
		b.SetCandidates(7, 9, []int{1, 2, 4, 7})
		b.SetCandidates(8, 9, []int{1, 2, 4})
		b.SetCandidates(9, 9, []int{2, 4, 6, 7})

		b2, e := NewBoard(dimensionSizeInBoxes, NewBox, NewCell)
		if b2 != nil {
			b2.SetCandidates(1, 1, []int{3, 5, 8})
			b2.SetValue(2, 1, 7)
			b2.SetValue(3, 1, 9)
			b2.SetCandidates(4, 1, []int{2, 3, 4, 8})
			b2.SetCandidates(5, 1, []int{2, 3, 4, 5})
			b2.SetCandidates(6, 1, []int{2, 8})
			b2.SetValue(7, 1, 6)
			b2.SetCandidates(8, 1, []int{2, 4, 5, 8})
			b2.SetValue(9, 1, 1)

			b2.SetCandidates(1, 2, []int{3, 5, 8})
			b2.SetCandidates(2, 2, []int{5, 8})
			b2.SetValue(3, 2, 1)
			b2.SetCandidates(4, 2, []int{2, 3, 4, 8})
			b2.SetValue(5, 2, 6)
			b2.SetValue(6, 2, 7)
			b2.SetCandidates(7, 2, []int{2, 4, 9})
			b2.SetCandidates(8, 2, []int{2, 4, 5, 8})
			b2.SetCandidates(9, 2, []int{2, 4, 5, 8, 9})

			b2.SetValue(1, 3, 6)
			b2.SetValue(2, 3, 2)
			b2.SetValue(3, 3, 4)
			b2.SetValue(4, 3, 9)
			b2.SetCandidates(5, 3, []int{1, 5})
			b2.SetCandidates(6, 3, []int{1, 8})
			b2.SetValue(7, 3, 3)
			b2.SetValue(8, 3, 7)
			b2.SetCandidates(9, 3, []int{5, 8})

			b2.SetCandidates(1, 4, []int{1, 5})
			b2.SetCandidates(2, 4, []int{1, 5})
			b2.SetValue(3, 4, 6)
			b2.SetCandidates(4, 4, []int{2, 3, 7, 8})
			b2.SetCandidates(5, 4, []int{2, 3, 7, 9})
			b2.SetValue(6, 4, 4)
			b2.SetCandidates(7, 4, []int{1, 2, 7, 9})
			b2.SetCandidates(8, 4, []int{1, 2, 3, 5, 8})
			b2.SetCandidates(9, 4, []int{2, 5, 7, 8, 9})

			b2.SetCandidates(1, 5, []int{4, 8})
			b2.SetCandidates(2, 5, []int{4, 8})
			b2.SetCandidates(3, 5, []int{2, 7})
			b2.SetValue(4, 5, 1)
			b2.SetCandidates(5, 5, []int{2, 3, 7, 9})
			b2.SetValue(6, 5, 5)
			b2.SetCandidates(7, 5, []int{2, 7, 9})
			b2.SetValue(8, 5, 6)
			b2.SetCandidates(9, 5, []int{2, 7, 9})

			b2.SetValue(1, 6, 9)
			b2.SetValue(2, 6, 3)
			b2.SetCandidates(3, 6, []int{2, 7})
			b2.SetCandidates(4, 6, []int{2, 6, 7, 8})
			b2.SetCandidates(5, 6, []int{2, 7})
			b2.SetCandidates(6, 6, []int{2, 6, 8})
			b2.SetCandidates(7, 6, []int{1, 2, 4, 7})
			b2.SetCandidates(8, 6, []int{1, 2, 4, 5, 8})
			b2.SetCandidates(9, 6, []int{2, 4, 5, 7, 8})

			b2.SetCandidates(1, 7, []int{1, 2, 4})
			b2.SetValue(2, 7, 6)
			b2.SetValue(3, 7, 5)
			b2.SetCandidates(4, 7, []int{2, 4, 7})
			b2.SetCandidates(5, 7, []int{1, 2, 4, 7})
			b2.SetCandidates(6, 7, []int{1, 2})
			b2.SetValue(7, 7, 8)
			b2.SetValue(8, 7, 9)
			b2.SetValue(9, 7, 3)

			b2.SetValue(1, 8, 7)
			b2.SetValue(2, 8, 9)
			b2.SetValue(3, 8, 8)
			b2.SetCandidates(4, 8, []int{2, 4, 6})
			b2.SetCandidates(5, 8, []int{1, 2, 4})
			b2.SetValue(6, 8, 3)
			b2.SetValue(7, 8, 5)
			b2.SetCandidates(8, 8, []int{1, 2, 4})
			b2.SetCandidates(9, 8, []int{2, 4, 6})

			b2.SetCandidates(1, 9, []int{1, 2, 4})
			b2.SetCandidates(2, 9, []int{1, 4})
			b2.SetValue(3, 9, 3)
			b2.SetValue(4, 9, 5)
			b2.SetValue(5, 9, 8)
			b2.SetValue(6, 9, 9)
			b2.SetCandidates(7, 9, []int{1, 2, 4, 7})
			b2.SetCandidates(8, 9, []int{1, 2, 4})
			b2.SetCandidates(9, 9, []int{2, 4, 6, 7})
			b.FindNakedPair(2, 5)
			b.PrintPossibilities()
			b2.PrintPossibilities()
			if !b.Equals(b2) {
				t.Errorf("2 boards are not equivalent")
			}

		} else {
			t.Errorf(e.Error())
		}
	} else {
		t.Errorf(e.Error())
	}
}
