import unittest
from typing import List
from board import Board
import puzzles

def lists_equal(l1, l2) -> bool:

    l1_tuple = []
    for r in l1:
        l1_tuple.append(tuple(r))

    t1 = tuple(l1_tuple)

    l2_tuple = []
    for r in l2:
        l2_tuple.append(tuple(r))
    t2 = tuple(l2_tuple)

    return t1 == t2

class TestBoard(unittest.TestCase):

    def test_nothing_set(self) -> None:
        board: Board = Board(size=3)
        c: int
        r: int
        for c in range(1, board.max_value + 1, 1):
            for r in range(1, board.max_value + 1, 1):
                self.assertEqual(board.get_value(c, r), -1)

    def test_value_set(self) -> None:
        board = Board(size=3)
        c_initial: int = 3
        r_initial: int = 5
        v_initial: int = 7
        board.set_cell(c_initial, r_initial, v_initial)

        c: int
        r: int
        for c in range(1, board.max_value + 1, 1):
            for r in range(1, board.max_value + 1, 1):
                value: int = board.get_value(c, r)
                if c == c_initial and r == r_initial:
                    self.assertEqual(v_initial, value)
                else:
                    self.assertEqual(-1, value)

    def test_get_representation(self) -> None:
        board: Board = Board(size=3)
        b: List[List[int]] = board.get_board_representation()
        self.assertEqual(board.max_value, len(b))
        row: List[int]
        for row in b:
            self.assertEqual(board.max_value, len(row))
            value: int
            for value in row:
                self.assertEqual(-1, value)

    def test_board_initialization(self) -> None:
        board: Board = Board(board=puzzles.pattern1_board)

        result_board: List[List[int]] = board.get_board_representation()
        self.assertTrue(lists_equal(puzzles.pattern1_board, result_board))

    def test_pattern(self) -> None:
        board: Board = Board(size=3)
        i: int
        for i in range(1, board.max_value + 1, 1):
            board.set_cell(i, i, 8)

        result_board = board.get_board_representation()

        self.assertTrue(lists_equal(puzzles.diagonal_eight_board, result_board))

    def test_valid_good_solution(self) -> None:
        board: Board = Board(board=puzzles.solved_board_1)
        self.assertTrue(board.valid())
        self.assertTrue(board.solved())

    def test_valid_partial_solution(self) -> None:
        board: Board = Board(board=puzzles.solvable_board_1)
        self.assertTrue(board.valid())

    def test_solve_cell(self) -> None:
        board: Board = Board(board=puzzles.solvable_board_1)

        # column 3, row 1 is solvable.
        self.assertTrue(board.solve_cell(3, 1))

        self.assertEqual(5, board.get_value(3, 1))

    def test_single_pass_solve(self) -> None:
        board: Board = Board(board=puzzles.solved_board_1)
        before: List[List[int]] = board.get_board_representation()

        # Puzzle is already solved, so no changes should occur.
        self.assertFalse(board.single_pass_solve())
        after: List[List[int]] = board.get_board_representation()
        self.assertTrue(lists_equal(before, after))

    def test_solve(self) -> None:
        board: Board = Board(board=puzzles.solvable_board_1)

        while board.single_pass_solve():
            pass

        solution: List[List[int]] = board.get_board_representation()
        self.assertTrue(lists_equal(solution, puzzles.solved_board_1))
        self.assertTrue(board.solved())

    def test_unsolvable(self) -> None:
        #Uninitialized board is all -1s
        board: Board = Board(size=3)
        while board.single_pass_solve():
            pass
        self.assertFalse(board.solved())

    def test_difficult_puzzle(self) -> None:

        board: Board = Board(board=puzzles.difficult_board)
        while board.single_pass_solve():
            pass
        board.print_board()
