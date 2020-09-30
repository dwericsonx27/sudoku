#include <set>

#include "sudokuCell.h"
#include "perfect.h"

SudokuCell::SudokuCell(int maxValue)
{
   setMaxValue(maxValue);
   setValue(-1);
}
/*
SudokuCell::~SudokuCell()
{
}
*/

void SudokuCell::setValue(int value)
{
}

void SudokuCell::setMaxValue(int value)
{
    if (value >= 1)
    {
        if (isPerfectSquare(value))
        {
        }
    }
    m_maxValue = value;
}

bool SudokuCell::eliminatePossiblies(const std::set<int>& eliminated)
{
    return true;
}
