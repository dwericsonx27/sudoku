"""Test the functions related to integer square roots."""

import unittest
import perfect

class TestPerfect(unittest.TestCase):
    """Test class for int_squre_root and is_perfect functions."""

    def test_int_square_root(self):
        """Test interesting values to determine the root."""
        self.assertEqual(2, perfect.int_square_root(4))
        self.assertEqual(1, perfect.int_square_root(1))
        self.assertEqual(0, perfect.int_square_root(0))
        self.assertEqual(2, perfect.int_square_root(5))
        self.assertEqual(1, perfect.int_square_root(3))
        with self.assertRaises(ValueError):
            perfect.int_square_root(-3)

    def test_is_perfect_square(self):
        """Testing interesting values for determine if a perfect square."""
        self.assertTrue(perfect.is_perfect_square(4))
        self.assertFalse(perfect.is_perfect_square(3))
        self.assertFalse(perfect.is_perfect_square(7))
        self.assertFalse(perfect.is_perfect_square(-7))
        self.assertTrue(perfect.is_perfect_square(1))
