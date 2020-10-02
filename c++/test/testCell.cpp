#include <gtest/gtest.h>
#include <sudokuCell.h>

TEST(CellTest, TestValueGreaterThanMaxValue){
  SudokuCell c(9);
  try {
    c.setValue(1001);
  }
  catch(std::out_of_range const & err) {
    EXPECT_EQ(err.what(),std::string("Value exceeds maximum value"));
  }
  catch(...) {
    FAIL() << "Expected std::out_of_range";
  }
}

TEST(CellTest, TestValueLessThanZero){
  SudokuCell c(9);
  try {
    c.setValue(-1);
  }
  catch(std::out_of_range const & err) {
    EXPECT_EQ(err.what(),std::string("Value must be positive"));
  }
  catch(...) {
    FAIL() << "Expected std::out_of_range";
  }
}

TEST(CellTest, TestValuePersisted){
  SudokuCell c(9);
  c.setValue(1);
  ASSERT_EQ(c.getValue(), 1);
}

TEST(CellTest, TestAllCandidates){
  SudokuCell c(9);
  const std::set<int>& candidates = c.getCandidates();
  ASSERT_EQ(candidates.size(), 9);
}

TEST(CellTest, TestRemoveCandidates){
  SudokuCell c(9);
  c.eliminatePossibilities({ 1, 2, 3 });
  const std::set<int>& candidates = c.getCandidates();
  ASSERT_EQ(candidates.size(), 6);
}

TEST(CellTest, TestNoCandidates){
  SudokuCell c(9);
  c.setValue(1);
  const std::set<int>& candidates = c.getCandidates();
  ASSERT_EQ(candidates.size(), 0);
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
