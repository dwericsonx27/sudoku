#include <cmath>

int intSquareRoot(int value)
{
    // Calculates the integer square root of an integer.

    return int(sqrt(value));
}

int isPerfectSquare(int value)
{
    int root = intSquareRoot(value);
    return (root * root == value);
}
