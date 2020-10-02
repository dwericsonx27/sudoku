#include <set>
#include <iostream>

#include "sudokuCell.h"
#include "perfect.h"
#include <stdexcept>

SudokuCell::SudokuCell(int maxValue)
{
  setMaxValue(maxValue);
  m_value = -1;
  for (int i = 1; i <= maxValue; ++i)
  {
    m_candidateValues.insert(i);
  }
}
/*
SudokuCell::~SudokuCell()
{
}
*/

void SudokuCell::setValue(int value)
{
  if (value > m_maxValue)
  {
    throw std::out_of_range("Value exceeds maximum value");
  }

  if (value < 0)
  {
    throw std::out_of_range("Value must be positive");
  }
  m_value = value;
  m_candidateValues.clear();
}

void SudokuCell::setMaxValue(int value)
{
  std::cerr << "setMaxValue(" << value << ")" << std::endl;
  if (isPerfectSquare(value))
  {
    m_maxValue = value;
    return;
  }
 
  throw std::out_of_range("Maximum value is not a perfect square");

}

int SudokuCell::getValue() const
{
    return m_value;
}

void SudokuCell::eliminatePossiblies(const std::set<int>& eliminated)
{
  for (auto it : eliminated)
  {
    m_candidateValues.erase(it);
  }
}
