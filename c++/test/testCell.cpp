#include <gtest/gtest.h>
#include <sudokuCell.h>

TEST(CellTest, TestNumberOne){
  SudokuCell c(9);
  EXPECT_NE(2, 1);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
