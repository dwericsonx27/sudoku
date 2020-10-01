#include <cmath>
#include <stdexcept>

int intSquareRoot(int value)
{
    // Calculates the integer square root of an integer.
    if (value < 0)
    {
        throw std::out_of_range("Out of range");
    }
    return int(sqrt(value));
}

int isPerfectSquare(int value)
{
    int root = intSquareRoot(value);
    return (root * root == value);
}
