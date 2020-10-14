#include "sudokuBox.h"
#include "perfect.h"

//        Box...
//   <--- dimension in Cells --->
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------
//   |        |        |        |
//   |  Cell  |  Cell  |  Cell  |
//   |        |        |        |
//   ----------------------------

SudokuBox::SudokuBox(int dimensionInCells, std::function<std::shared_ptr<Cell>()> cellCreator)
{
  m_dimension = dimensionInCells;
  int cellCount = dimensionInCells * dimensionsInCells;
  for (int i = 0; i < cellCount; ++i)
  {
    m_box[i] = cellCreator();
  }
}

shared_ptr<Cell> SudokuBox::getCell(int column, int row)
{
  int cellNum = (m_dimension * row) + column;
  return m_box[cellNum];
}

bool SudokuBox::isValid()
{
}
