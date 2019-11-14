//
// mix of simple array examples
//

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func pivotIndex(nums []int) int {
	var s1, s2 int

	if len(nums) < 3 {
		return -1
	}

	for _, v := range nums {
		s2 += v
	}

	for i, v := range nums {
		if s1 == s2-v {
			return i
		}

		s1 += v
		s2 -= v
	}

	return -1
}

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
