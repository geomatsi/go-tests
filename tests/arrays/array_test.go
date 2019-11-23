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
