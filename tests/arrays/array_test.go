//
// mix of simple array examples
//

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, pivotIndex([]int{}), -1)
	assert.Equal(t, pivotIndex([]int{1}), -1)
	assert.Equal(t, pivotIndex([]int{1, 2}), -1)
	assert.Equal(t, pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
	assert.Equal(t, pivotIndex([]int{1, 1, 1, 1, 1, 1, 1}), 3)
	assert.Equal(t, pivotIndex([]int{-7, 1, 5, 2, -4, 3, 0}), 3)
	assert.Equal(t, pivotIndex([]int{-1, -1, -1, 0, 1, 1}), 0)
	assert.Equal(t, pivotIndex([]int{-1, -1, 0, 1, 1, 0}), 5)
	assert.Equal(t, pivotIndex([]int{1, 2, 3}), -1)
	assert.Equal(t, pivotIndex([]int{1, 2, 3, 4}), -1)
}

func TestDominantIndex(t *testing.T) {
	assert.Equal(t, dominantIndex([]int{}), -1)
	assert.Equal(t, dominantIndex([]int{2, 3}), -1)
	assert.Equal(t, dominantIndex([]int{3, 2, 1, 0}), -1)
	assert.Equal(t, dominantIndex([]int{3, 3, 1, 5}), -1)
	assert.Equal(t, dominantIndex([]int{1}), 0)
	assert.Equal(t, dominantIndex([]int{1, 2}), 1)
	assert.Equal(t, dominantIndex([]int{3, 6, 1, 0}), 1)
	assert.Equal(t, dominantIndex([]int{6, 3, 1, 0}), 0)
	assert.Equal(t, dominantIndex([]int{1, 3, 1, 7}), 3)

}

func TestPlusOne(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, plusOne([]int{0}), []int{1})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
	assert.Equal(t, plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})
	assert.Equal(t, plusOne([]int{4, 3, 3, 1}), []int{4, 3, 3, 2})
}

func TestFindDiagOrder(t *testing.T) {
	assert.Equal(t, findDiagonalOrder([][]int{}), []int{})
	assert.Equal(t, findDiagonalOrder([][]int{{}}), []int{})
	assert.Equal(t, findDiagonalOrder([][]int{{}, {}}), []int{})
	assert.Equal(t, findDiagonalOrder([][]int{{1}}), []int{1})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2}}), []int{1, 2})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2, 3, 4}}), []int{1, 2, 3, 4})
	assert.Equal(t, findDiagonalOrder([][]int{{1}, {2}, {3}, {4}}), []int{1, 2, 3, 4})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4})
	assert.Equal(t, findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 4, 7, 5, 3, 6, 8, 9})
}

func TestSpiralOrder(t *testing.T) {
	assert.Equal(t, spiralOrder([][]int{}), []int{})
	assert.Equal(t, spiralOrder([][]int{{}}), []int{})
	assert.Equal(t, spiralOrder([][]int{{}, {}}), []int{})
	assert.Equal(t, spiralOrder([][]int{{1}}), []int{1})
	assert.Equal(t, spiralOrder([][]int{{1, 2}}), []int{1, 2})
	assert.Equal(t, spiralOrder([][]int{{1, 2, 3, 4}}), []int{1, 2, 3, 4})
	assert.Equal(t, spiralOrder([][]int{{1}, {2}, {3}, {4}}), []int{1, 2, 3, 4})
	assert.Equal(t, spiralOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 4, 3})
	assert.Equal(t, spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	assert.Equal(t, spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}

func TestPascalTriangle(t *testing.T) {
	assert.Equal(t, generate(0), [][]int{})
	assert.Equal(t, generate(1), [][]int{{1}})
	assert.Equal(t, generate(2), [][]int{{1}, {1, 1}})
	assert.Equal(t, generate(3), [][]int{{1}, {1, 1}, {1, 2, 1}})
	assert.Equal(t, generate(4), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}})
	assert.Equal(t, generate(5), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
	assert.Equal(t, generate(6), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}})
	assert.Equal(t, generate(7), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}})
}

func TestArrayPairSum(t *testing.T) {
	assert.Equal(t, 4, arrayPairSum([]int{1, 4, 3, 2}))
	assert.Equal(t, 3, arrayPairSum([]int{1, 1, 2, 10}))
	assert.Equal(t, 3, arrayPairSum([]int{1, 1, 2, 10}))
	assert.Equal(t, 9, arrayPairSum([]int{1, 2, 3, 4, 5, 6}))
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{2, 4}, twoSum([]int{1, 2, 3, 4, 6}, 6))
	assert.Equal(t, []int{2, 3}, twoSum([]int{1, 10, 21, 22, 31, 40}, 31))
}

func TestRemoveElements(t *testing.T) {
	var num []int
	var len int

	num = []int{3, 2, 2, 3}
	len = removeElement(num, 3)
	assert.Equal(t, len, 2)
	assert.Equal(t, []int{2, 2}, num[0:len])

	num = []int{2, 3, 2, 3}
	len = removeElement(num, 3)
	assert.Equal(t, len, 2)
	assert.Equal(t, []int{2, 2}, num[0:len])

	num = []int{2, 2, 3, 3}
	len = removeElement(num, 3)
	assert.Equal(t, len, 2)
	assert.Equal(t, []int{2, 2}, num[0:len])

	num = []int{0, 1, 2, 2, 3, 0, 4, 2}
	len = removeElement(num, 2)
	assert.Equal(t, len, 5)
	assert.Equal(t, []int{0, 1, 3, 0, 4}, num[0:len])
}

func TestFindMaxConsOnes(t *testing.T) {
	assert.Equal(t, 0, findMaxConsecutiveOnes([]int{}))
	assert.Equal(t, 0, findMaxConsecutiveOnes([]int{0}))
	assert.Equal(t, 1, findMaxConsecutiveOnes([]int{1}))
	assert.Equal(t, 1, findMaxConsecutiveOnes([]int{1, 0}))
	assert.Equal(t, 2, findMaxConsecutiveOnes([]int{1, 1}))
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 1, 0, 1, 1}))
	assert.Equal(t, 5, findMaxConsecutiveOnes([]int{0, 1, 1, 1, 1, 1, 0}))
}

func TestMinSubArrayLen(t *testing.T) {
	assert.Equal(t, 0, minSubArrayLen(1, []int{}))
	assert.Equal(t, 0, minSubArrayLen(1, []int{0}))
	assert.Equal(t, 1, minSubArrayLen(1, []int{1}))
	assert.Equal(t, 1, minSubArrayLen(1, []int{1, 0, 1, 0}))
	assert.Equal(t, 3, minSubArrayLen(2, []int{1, 0, 1, 0}))
	assert.Equal(t, 4, minSubArrayLen(4, []int{1, 1, 1, 1}))
	assert.Equal(t, 3, minSubArrayLen(4, []int{2, 1, 1, 0, 0, 1}))
	assert.Equal(t, 2, minSubArrayLen(4, []int{2, 1, 1, 0, 0, 1, 3}))
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func TestRotate(t *testing.T) {
	var nums []int

	nums = []int{}
	rotate(nums, 1)
	assert.Equal(t, nums, []int{})

	nums = []int{1}
	rotate(nums, 5)
	assert.Equal(t, []int{1}, nums)

	nums = []int{1, 2}
	rotate(nums, 0)
	assert.Equal(t, []int{1, 2}, nums)

	nums = []int{1, 2}
	rotate(nums, 1)
	assert.Equal(t, []int{2, 1}, nums)

	nums = []int{1, 2}
	rotate(nums, 2)
	assert.Equal(t, []int{1, 2}, nums)

	nums = []int{1, 2, 3, 4}
	rotate(nums, 1)
	assert.Equal(t, []int{4, 1, 2, 3}, nums)

	nums = []int{1, 2, 3, 4}
	rotate(nums, 2)
	assert.Equal(t, []int{3, 4, 1, 2}, nums)
}

func TestRotate1(t *testing.T) {
	var nums []int

	nums = []int{}
	rotate1(nums, 1)
	assert.Equal(t, nums, []int{})

	nums = []int{1}
	rotate1(nums, 5)
	assert.Equal(t, []int{1}, nums)

	nums = []int{1, 2}
	rotate1(nums, 0)
	assert.Equal(t, []int{1, 2}, nums)

	nums = []int{1, 2}
	rotate1(nums, 1)
	assert.Equal(t, []int{2, 1}, nums)

	nums = []int{1, 2}
	rotate1(nums, 2)
	assert.Equal(t, []int{1, 2}, nums)

	nums = []int{1, 2, 3, 4}
	rotate1(nums, 1)
	assert.Equal(t, []int{4, 1, 2, 3}, nums)

	nums = []int{1, 2, 3, 4}
	rotate1(nums, 2)
	assert.Equal(t, []int{3, 4, 1, 2}, nums)
}

func TestGetPascalRow(t *testing.T) {
	assert.Equal(t, []int{1}, getPascalRow(0))
	assert.Equal(t, []int{1, 1}, getPascalRow(1))
	assert.Equal(t, []int{1, 2, 1}, getPascalRow(2))
	assert.Equal(t, []int{1, 3, 3, 1}, getPascalRow(3))
	assert.Equal(t, []int{1, 4, 6, 4, 1}, getPascalRow(4))
	assert.Equal(t, []int{1, 5, 10, 10, 5, 1}, getPascalRow(5))
}

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 0, removeDuplicates([]int{}))
	assert.Equal(t, 1, removeDuplicates([]int{1}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 2}))
	assert.Equal(t, 1, removeDuplicates([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, 4, removeDuplicates([]int{0, 1, 2, 3}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func TestMoveZeroes(t *testing.T) {
	var nums []int

	nums = []int{}
	moveZeroes(nums)
	assert.Equal(t, []int{}, nums)

	nums = []int{1}
	moveZeroes(nums)
	assert.Equal(t, []int{1}, nums)

	nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, nums)

	nums = []int{1, 0, 0, 0, 1, 0, 0, 1, 1, 0}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0}, nums)
}
