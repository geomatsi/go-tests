//
// mix of simple array examples
//

package main

import (
	"fmt"
)

func main() {
	fmt.Println("run tests: go test -cover")
}

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

func findDiagonalOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}

	m := len(matrix[0])
	if m == 0 {
		return []int{}
	}

	res, ind := make([]int, n*m), 0

	for k := 0; k < (n + m - 1); k++ {
		for i := 0; i < (k + 1); i++ {
			if k%2 == 0 {
				if i < m && (k-i) < n {
					res[ind], ind = matrix[k-i][i], ind+1
				}
			} else if i < n && (k-i) < m {
				res[ind], ind = matrix[i][k-i], ind+1
			}
		}
	}

	return res
}
