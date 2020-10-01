#include <gtest/gtest.h>
#include <perfect.h>

TEST(PerfectTest, TestIntSquareRoot) {
  ASSERT_EQ(2, intSquareRoot(4));
  ASSERT_EQ(1, intSquareRoot(1));
  ASSERT_EQ(0, intSquareRoot(0));
  ASSERT_EQ(2, intSquareRoot(5));
  ASSERT_EQ(1, intSquareRoot(3));
  try {
    ASSERT_EQ(1, intSquareRoot(-3));
  }
  catch(std::out_of_range const & err) {
    EXPECT_EQ(err.what(),std::string("Out of range"));
  }
  catch(...) {
    FAIL() << "Expected std::out_of_range";
  }
}

TEST(PerfectTest, TestPerfectSquare){
  ASSERT_TRUE(isPerfectSquare(4));
  ASSERT_FALSE(isPerfectSquare(3));
  ASSERT_FALSE(isPerfectSquare(7));
  try {
    ASSERT_FALSE(isPerfectSquare(-7));
  }
  catch(std::out_of_range const & err) {
    EXPECT_EQ(err.what(),std::string("Out of range"));
  }
  catch(...) {
    FAIL() << "Expected std::out_of_range";
  }
  ASSERT_TRUE(isPerfectSquare(1));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
