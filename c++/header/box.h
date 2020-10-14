#ifndef ABSTRACT_BOX_H
#define ABSTRACT_BOX_H

class Box {
    public:
        virtual shared_ptr<Cell> getCell(int column, int row) = 0;
	virtual bool isValid() = 0;
};

#endif
