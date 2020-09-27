from typing import Dict
from cell import Cell

#        Box...
#   <--- dimension in Cells --->
#   ----------------------------
#   |        |        |        |
#   |  Cell  |  Cell  |  Cell  |
#   |        |        |        |
#   ----------------------------
#   |        |        |        |
#   |  Cell  |  Cell  |  Cell  |
#   |        |        |        |
#   ----------------------------
#   |        |        |        |
#   |  Cell  |  Cell  |  Cell  |
#   |        |        |        |
#   ----------------------------

class Box:

    def __init__(self, dimensionInCells: int):
        self.cells: Dict[int, Cell] = {}
        self.dimensionInCells: int = dimensionInCells
        self.max_value: int = dimensionInCells * dimensionInCells
        i: int = 0
        for i in range(1, self.max_value + 1, 1):
            self.cells[i] = Cell(self.max_value)

    def __str__(self):
        rtnval: str = 'Box:\n'
        i: int = 0
        for i in range(1, self.max_value + 1, 1):
            rtnval += f'{self.cells[i]}'
        return rtnval

    def get_cell(self, column: int, row: int):
        cell_num: int = column + (row-1) * self.dimensionInCells
        return self.cells[cell_num]

    def is_valid(self):
        findings: Dict[int, int] = {}
        c: int
        r: int
        for c in range(1, self.dimensionInCells + 1, 1):
            for r in range(1, self.dimensionInCells + 1, 1):
                value: int = self.get_cell(c, r).value
                if value > 0:
                    if value in findings:
                        return False
                    else:
                        findings[value] = 1
        return True
