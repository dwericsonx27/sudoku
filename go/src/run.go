package main

import (
	"fmt"
	"sudoku"
)

var solvableBoard1 [][]int = [][]int{
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

func testNakedPair() {
    dimensionSizeInBoxes := 3
    b, e := sudoku.NewBoard(dimensionSizeInBoxes, sudoku.NewBox, sudoku.NewCell)
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
	b.PrintPossibilities()
        if b.FindNakedPair(2, 5) {
            b.PrintPossibilities()
        } else {
            fmt.Printf("No Changes Made!")
        }
    } else {
	fmt.Printf("ERROR: %s\n", e.Error())
    }
}

func main() {
	b, e := sudoku.NewBoardInitialize(solvableBoard1)
	if b != nil {
		b.Print()
		b.PrintPossibilities()
	} else {
		fmt.Printf("ERROR: %s\n", e.Error())
	}
	ok, err := b.IsValid()
	if !ok {
		fmt.Printf("ERROR: board invalid!\n")
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
	b.Solve()
	b.Print()
	b.PrintPossibilities()
        testNakedPair()
}
