from board import Board
import puzzles

b = Board(board=puzzles.solvable_board_1)

b.print_board()
print('--------------------------------------------------')
b.print_possibilities()
print('--------------------------------------------------')
while b.single_pass_solve():
    pass
b.print_board()

b = Board(board=puzzles.difficult_board)
b.print_board()
print("")
while b.single_pass_solve():
    pass
b.print_possibilities()

b = Board(board=puzzles.difficult_board_correct_change)
b.print_board()
print("")
while b.single_pass_solve():
    pass
b.print_possibilities()

b = Board(board=puzzles.difficult_board_wrong_change)
b.print_board()
print("")
while b.single_pass_solve():
    pass
b.print_possibilities()
