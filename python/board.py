from typing import Dict, Set, Tuple, Any, List
from cell import Cell
from box import Box
import perfect

#          Board...
#   <--- dimension in Boxes --->
#   ----------------------------
#   |        |        |        |
#   |  Box   |  Box   |  Box   |
#   |        |        |        |
#   ----------------------------
#   |        |        |        |
#   |  Box   |  Box   |  Box   |
#   |        |        |        |
#   ----------------------------
#   |        |        |        |
#   |  Box   |  Box   |  Box   |
#   |        |        |        |
#   ----------------------------

class Board:

    def __init__(self,  **kwargs: Any):
        self.boxes: Dict[int, Box] = {}
        if 'board' in kwargs:
            representation=kwargs['board']
            self.max_value = len(representation)   
            if perfect.is_perfect_square(self.max_value):
                self.dimensionInBoxes = perfect.int_square_root(self.max_value)
                self._create_boxes()
                self._initialize_board(representation)
            else:
                raise RuntimeError(f'Board size is not a perfect square!')
        elif 'size' in kwargs:
            self.dimensionInBoxes: int = int(kwargs['size'])
            self.max_value: int = self.dimensionInBoxes * self.dimensionInBoxes
            self._create_boxes()
        else:
            raise RuntimeError(f'Board size not supplied!')

    def _create_boxes(self) -> None:
        print(f'Creating {self.max_value} boxes.')
        i: int
        for i in range(1, self.max_value + 1, 1):
            self.boxes[i] = Box(self.dimensionInBoxes)

    def __str__(self):
        rtnval: str = 'Board:'
        i: int
        for i in range(1, self.max_value + 1, 1):
            rtnval += f'{self.boxes[i]}'
        return rtnval

    def _column_row_to_boxnum(self, column: int, row: int) -> Tuple[int, int, int]:
        base_column: int = (column + 2)//self.dimensionInBoxes
        offset: int = self.dimensionInBoxes * ((row - 1)//self.dimensionInBoxes)
        box_num: int = base_column + offset
        box_c: int = 1 + ((column - 1)%self.dimensionInBoxes)
        box_r: int = 1 + ((row - 1)%self.dimensionInBoxes)
        return box_num, box_c, box_r

    def _get_cell(self, column: int, row: int) -> Cell:
        box_num, box_c, box_r = self._column_row_to_boxnum(column, row)
        cell = self.boxes[box_num].get_cell(box_c, box_r)
        return cell

    def _initialize_board(self, representation) -> None:
        self.max_value = len(representation)
        if perfect.is_perfect_square(self.max_value):
            self.dimensionInBoxes = perfect.int_square_root(self.max_value)
            r: int = 0
            for row in representation:
                r += 1
                assert(len(row) == self.max_value)
                c: int = 0
                for val in row:
                    c += 1
                    self.set_cell(c, r, val)
            return
        raise RuntimeError(f'Board size is not a perfect square!')

    def set_cell(self, column: int, row: int, value: int):
        cell = self._get_cell(column, row)
        cell.set_value(value)

    def get_value(self, column: int, row: int):
        return self._get_cell(column, row).value

    def _get_row(self, row):
        row_list_of_cells: List[Cell] = []
        for c in range(1, self.max_value + 1, 1):
            row_list_of_cells.append(self._get_cell(c, row))
        return row_list_of_cells

    def get_board_representation(self):
        rtnval = []
        for r in range(1, self.max_value + 1, 1):
            subject_row = []
            for c in range(1, self.max_value + 1, 1):
                subject_row.append(self.get_value(c, r))
            rtnval.append(subject_row)
        return rtnval

    def valid_column(self, col_num: int) -> bool:
       catalog: Dict[int, int] = {}
       for r in range(1, self.max_value + 1, 1):
           val = self.get_value(col_num, r)
           if val > 0:
               if val in catalog:
                   catalog[val] = catalog[val] + 1
               else:
                   catalog[val] = 1

       rtnval: bool = True
       for k, v in catalog.items():
           if v > 1:
               rtnval = False
               print(f'Found {k} {v} times in column {col_num}!')
       return rtnval

    def are_valid_columns(self) -> bool:
        rtnval: bool = True
        for i in range(1, self.max_value + 1, 1):
            rtnval = rtnval & self.valid_column(i) 
        return rtnval

    def valid_row(self, row_num: int) -> bool:
       catalog: Dict[int, int] = {}
       for c in range(1, self.max_value + 1, 1):
           val = self.get_value(c, row_num)
           if val > 0:
               if val in catalog:
                   catalog[val] = catalog[val] + 1
               else:
                   catalog[val] = 1

       rtnval: bool = True
       for k, v in catalog.items():
           if v > 1:
               rtnval = False
               print(f'Found {k} {v} times in row {row_num}!')
       return rtnval

    def are_valid_rows(self) -> bool:
        rtnval: bool = True
        for i in range(1, self.max_value + 1, 1):
            rtnval = rtnval & self.valid_row(i) 
        return rtnval

    def are_valid_boxes(self) -> bool:
        rtnval: bool = True
        for i in range(1, self.max_value + 1, 1):
            rtnval = rtnval & self.boxes[i].is_valid()
        return rtnval

    def valid(self) -> bool:
        ok_rows = self.are_valid_rows()
        ok_columns = self.are_valid_columns()
        ok_boxes = self.are_valid_boxes()

        return ok_rows & ok_columns & ok_boxes

    def _all_cells_determined(self) -> bool:
        for c in range(1, self.max_value + 1, 1):
            for r in range(1, self.max_value + 1, 1):
                if self.get_value(c, r) == -1:
                    return False
        return True
                
    def solved(self) -> bool:
        valid = self.valid()
        all_cells_determined: bool = self._all_cells_determined()

        return valid & all_cells_determined

    def find_values_column(self, column: int, used_values) -> None:
        for r in range(1, self.max_value + 1, 1):
            value = self.get_value(column, r)
            if 1 <= value <= self.max_value:
                used_values.add(self.get_value(column, r))

    def find_values_row(self, row: int, used_values) -> None:
        for c in range(1, self.max_value + 1, 1):
            value = self.get_value(c, row)
            if 1 <= value <= self.max_value:
                used_values.add(self.get_value(c, row))

    def find_values_box(self, column: int, row: int, used_values) -> None:
        box_num, _, _ = self._column_row_to_boxnum(column, row)
        box = self.boxes[box_num]
        for i in range(1, self.max_value + 1, 1):
            value = box.cells[i].value
            if 1 <= value <= self.max_value:
                used_values.add(value)

    def solve_cell(self, column: int, row: int) -> bool:
        cell = self._get_cell(column, row)
        if 1 <= cell.value <= self.max_value:
            return False

        used_values: Set[int] = set()
        self.find_values_column(column, used_values)
        self.find_values_row(row, used_values)
        self.find_values_box(column, row, used_values)

        return cell.discard_and_set_value(used_values)

    def single_pass_solve(self) -> bool:
        rtnval: bool = False
        for c in range(1, self.max_value + 1, 1):
            for r in range(1, self.max_value + 1, 1):
                status: bool = self.solve_cell(c, r)
                rtnval = rtnval | status
        return rtnval

    def print_board(self):
        print("")
        print('----------------------------')
        for r in range(1, self.max_value + 1, 1):
            for c in range(1, self.max_value + 1, 1):
                print('|', end="")
                print(f'{self._get_cell(c, r).value:2}', end="")
            print('|')
        print('----------------------------')

    def print_possibilities(self):
        row_divider: str = '-------------------------------------'
        print(row_divider)
        r: int
        for r in range(1, self.max_value + 1, 1):
            row_of_cells = self._get_row(r)
            for index in range(1, self.max_value, self.dimensionInBoxes):
                s: str = '|'
                for c in row_of_cells:
                    if 1 <= c.value <= self.max_value:
                        for _ in range(0, self.dimensionInBoxes, 1):
                            s = s + str(c.value)
                    else:
                        for offset in range(0, self.dimensionInBoxes, 1):
                            val: int = index + offset
                            if val in c.possibilities:
                                s = s + str(val)
                            else:
                                s = s + '.'
                    s = s + '|'
                print(s)
            print(row_divider)
