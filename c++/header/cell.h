#ifndef ABSTRACT_CELL_H
#define ABSTRACT_CELL_H

#include <set>

class Cell {
    public:
        virtual void setValue(int value) = 0;
	virtual void setMaxValue(int value) = 0;
        virtual bool eliminatePossiblies(const std::set<int>& eliminated) = 0;
    private:
};

#endif