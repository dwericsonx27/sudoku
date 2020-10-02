#ifndef SUDOKU_CELL_H
#define SUDOKU_CELL_H

#include <set>
#include "cell.h"

class SudokuCell : public Cell {
    public:
        SudokuCell(int value);
        void setValue(int value);
	void setMaxValue(int value);
        void eliminatePossibilities(const std::set<int>& eliminated);
        int getValue() const;
        const std::set<int>& getCandidates() const;
    private:
        int m_maxValue;
        int m_value;
        std::set<int> m_candidateValues;
};

#endif
