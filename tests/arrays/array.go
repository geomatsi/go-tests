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

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}

	m := len(matrix[0])
	if m == 0 {
		return []int{}
	}

	res, ind := make([]int, n*m), 0

	i1, i2, j1, j2 := 0, m-1, 0, n-1

	for {
		if i1 > i2 {
			break
		}

		for s := i1; s <= i2; s++ {
			res[ind], ind = matrix[j1][s], ind+1
		}

		if (j1 + 1) > j2 {
			break
		}

		for s := j1 + 1; s <= j2; s++ {
			res[ind], ind = matrix[s][i2], ind+1
		}

		if i1 == i2 {
			break
		}

		for s := i2 - 1; s >= i1; s-- {
			res[ind], ind = matrix[j2][s], ind+1
		}

		if (j1 + 1) == j2 {
			break
		}

		for s := (j2 - 1); s > j1; s-- {
			res[ind], ind = matrix[s][i1], ind+1
		}

		i1, i2, j1, j2 = i1+1, i2-1, j1+1, j2-1
	}

	return res
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)

		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}
