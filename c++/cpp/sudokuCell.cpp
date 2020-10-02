#include <set>
#include <iostream>
#include <stdexcept>

#include "sudokuCell.h"
#include "perfect.h"

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
  if (isPerfectSquare(value))
  {
    m_maxValue = value;
    return;
  }
 
  throw std::out_of_range("Maximum value is not a perfect square");

}

void SudokuCell::eliminatePossibilities(const std::set<int>& eliminated)
{
  // Remove the specified candidates.
  for (auto it : eliminated)
  {
    m_candidateValues.erase(it);
  }

  // When there are no candidates left directly after eliminating candidates,
  // there is a logic error since 1 candidate must be left.
  if (m_candidateValues.size() == 0)
  {
    throw std::runtime_error("Eliminated all possiblities, at least one must be left.");
  }

  // When 1 candidate is left, set the cell value and empty the candidate list.
  if (m_candidateValues.size() == 1)
  {
    auto it = m_candidateValues.begin();
    m_value = *it;
    m_candidateValues.clear();
  }
}

int SudokuCell::getValue() const
{
    return m_value;
}

const std::set<int>& SudokuCell::getCandidates() const
{
  return m_candidateValues;
}
