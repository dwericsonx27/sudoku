#ifndef SUDOKU_CELL_H
#define SUDOKU_CELL_H

#include "cell.h"

class SudokuCell : public Cell {
    public:
        SudokuCell(int value);
        void setValue(int value);
	void setMaxValue(int value);
        bool eliminatePossiblies(const std::set<int>& eliminated);
    private:
        int m_maxValue;
};

#endif
