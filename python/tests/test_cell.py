import unittest
from cell import Cell

class TestCell(unittest.TestCase):

    def test_set_valid_value(self):
        self.cell = Cell(9)
        value: int = 1
        self.cell.set_value(value)
        self.assertEqual(value, self.cell.value)

        #if the value is valid, then there should be no possiblities.
        self.assertEqual(0, len(self.cell.possibilities))

    def test_invalid_max_value(self):
        with self.assertRaises(RuntimeError): Cell(17)

    def test_set_unknown_value(self):
        self.cell = Cell(9)
        value: int = -1
        self.cell.set_value(value)
        self.assertEqual(value, self.cell.value)

        # All values are possible.
        self.assertEqual(9, len(self.cell.possibilities))

    def test_bad_value(self):
        self.cell = Cell(9)
        value: int = 100
        with self.assertRaises(RuntimeError): self.cell.set_value(value)
