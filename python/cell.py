"""Module to represent a single value/element on the suduko board."""
from typing import Set

import perfect

class Cell:
    """Class to represent a single value/element on the suduko board."""

    def __init__(self, max_value: int):
        self.max_value: int
        self.value: int
        # Setting max value must come first, so when the value is
        # set it can be evaluated.
        self.set_max_value(max_value)
        self.set_value(-1)

    def __str__(self):
        rtnval: str = 'Cell:\n'
        rtnval += f'      value: {self.value}\n'
        rtnval += f'      possibilities: {self.possibilities}\n'
        rtnval += f'      -----------------------------------\n'
        return rtnval

    def set_value(self, value: int) -> None:
        """Set the value for the cell, and validate it reasonable."""
        if not (value == -1 or 1 <= value <= self.max_value):
            raise RuntimeError(f'Invalid value for Cell: {value}')
        self.value: int = value
        self.possibilities: Set[int] = set()
        if value == -1:
            i: int
            for i in range(1, self.max_value + 1, 1):
                self.possibilities.add(i)

    def set_max_value(self, max_value: int) -> None:
        """Set what the maximum allowable value for a cell is."""
        if max_value >= 1:
            if perfect.is_perfect_square(max_value):
                self.max_value = max_value
                return
        raise RuntimeError(f'Invalid max value for Cell: {max_value}')

    def discard_and_set_value(self, used_values: Set[int]) -> bool:
        """Discard values in possibilites for the cell, if only one
           possiblity is left, assign that as the cell value."""
        if 1 <= self.value <= self.max_value:
            if self.possibilities:
                print(f'ERROR: value: {self.value}, possibilities: {self.possibilities}')
        value: int
        for value in used_values:
            self.possibilities.discard(value)
        if len(self.possibilities) == 1:
            self.value = self.possibilities.pop()
            return True
        return False
