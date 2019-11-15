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

// Assumptions:
// - nums will have a length in the range [1, 50]
// - every nums[i] will be an integer in the range [0, 99]
func dominantIndex(nums []int) int {
	var p1, v2 int

	switch {
	case len(nums) == 0:
		return -1
	case len(nums) == 1:
		return 0
	}

	for idx, val := range nums[1:] {
		if val > nums[p1] {
			if nums[p1] > v2 {
				v2 = nums[p1]
			}

			p1 = idx + 1
		} else if val > v2 {
			v2 = val
		}
	}

	if 2*v2 <= nums[p1] {
		return p1
	}

	return -1
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

func plusOne(digits []int) []int {
	r := len(digits)
	z := 0

	digits[r-1]++

	for i := r - 1; i >= 0; i-- {
		digits[i] += z
		if digits[i] >= 10 {
			digits[i] -= 10
			z = 1
		} else {
			z = 0
			break
		}
	}

	if z == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func TestPlusOne(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, plusOne([]int{0}), []int{1})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
	assert.Equal(t, plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})
	assert.Equal(t, plusOne([]int{4, 3, 3, 1}), []int{4, 3, 3, 2})
}
