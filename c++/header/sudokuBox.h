#ifndef SUDOKU_BOX_H
#define SUDOKU_BOX_H

#include "box.h"
#include "cell.h"

#include <vector>

class SudokuBox : public Box {
  public:
    SudokuBox(int dimensionInCells, std::function<std::shared_ptr<Cell>()> cellCreator);
    virtual shared_ptr<Cell> getCell(int column, int row) = 0;
    virtual bool isValid() = 0;
  private:
    std::vector<std::shared_ptr<Cell>> m_box;
    int m_dimension;
};

#endif
