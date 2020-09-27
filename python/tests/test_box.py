import unittest
from box import Box

class TestBox(unittest.TestCase):

    def setUp(self) -> None:
        self.box = Box(3)

    def test_initialization(self) -> None:
        for c in range(1, 4, 1):       
            for r in range(1, 4, 1):       
                self.assertEqual(-1, self.box.get_cell(c, r).value)

    def test_get_cell(self) -> None:
       cell = self.box.get_cell(2, 2)
       value: int = 5
       cell.set_value(value)

    def test_validation_initial(self) -> None:
        self.assertTrue(self.box.is_valid())

    def test_validation_valid(self) -> None:
        for c in range(1, 4, 1):
            for r in range(1, 4, 1):
                cell = self.box.get_cell(c, r)
                cell.set_value(3*(r-1) + c)

        self.assertTrue(self.box.is_valid())

    def test_validation_invalid(self) -> None:
        for c in range(1, 4, 1):
            for r in range(1, 4, 1):
                cell = self.box.get_cell(c, r)
                cell.set_value(c)

        self.assertFalse(self.box.is_valid())

                
