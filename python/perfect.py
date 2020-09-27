import math

def int_square_root(value: int) -> int:
    """Calculates the integer square of an integer.

       Attempted to write on function using Newton's method however
       it ran at half the speed of this metod."""
    return int(math.sqrt(value))

def is_perfect_square(n: int) -> int:
    try:
        root: int = int_square_root(n)
        return root*root == n
    except ValueError:
        return False
