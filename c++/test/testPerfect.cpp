#include <gtest/gtest.h>
#include <perfect.h>

TEST(PerfectTest, TestIntSquareRoot){
  ASSERT_EQ(3, intSquareRoot(9));
  ASSERT_NE(3, intSquareRoot(8));
}

TEST(PerfectTest, TestPerfectSquare){
  ASSERT_TRUE(isPerfectSquare(9));
  ASSERT_FALSE(isPerfectSquare(8));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
