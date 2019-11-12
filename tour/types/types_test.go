//
// mix of examples from 'A Tour of Go': section 'Flow control statements'
//

package main

import (
	"strings"
	"testing"
)

func TestPointers(t *testing.T) {
	var i int = 42
	var p *int = &i

	if *p != 42 {
		t.Errorf("expected 42, got %v", *p)
	}

	i = 2

	if *p != 2 {
		t.Errorf("expected 2, got %v", *p)
	}

	*p = 3

	if i != 3 {
		t.Errorf("expected 3, got %v", i)
	}
}

type TestStruct struct {
	x int
	y int
}

func TestStructs(t *testing.T) {
	var s TestStruct = TestStruct{4, 2}

	if s.x != 4 || s.y != 2 {
		t.Errorf("unexpected struct fields: (%v, %v) != (4, 2)", s.x, s.y)
	}

	s.x += 1

	if s.x != 5 || s.y != 2 {
		t.Errorf("unexpected struct fields: (%v, %v) != (5, 2)", s.x, s.y)
	}

	var p *TestStruct = &s

	if p.x != s.x || (*p).x != s.x {
		t.Errorf("struct pointer access error: %v, %v, %v)", s.x, p.x, (*p).x)
	}

	var sp TestStruct = TestStruct{x: 4}

	if !(sp.x == 4 && sp.y == 0) {
		t.Errorf("unexpected struct fields: (%v, %v) != (4, 0)", sp.x, sp.y)
	}

}

func TestArrays(t *testing.T) {
	var a [5]int

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	b := [5]int{1, 2, 3, 4, 5}

	if a != b {
		t.Errorf("arrays are not equal: %v != %v", a, b)
	}

	// slice for idx in [1, 3)
	var s1 []int = a[1:3]

	if len(s1) != 2 {
		t.Errorf("incorrect slice len: %v -> %v", s1, len(s1))
	}

	// the capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice

	if cap(s1) != 4 {
		t.Errorf("incorrect slice cap: %v -> %v", s1, cap(s1))
	}

	if s1[0] != 2 || s1[1] != 3 {
		t.Errorf("incorrect slice: %v -> %v", a, s1)
	}

	// modify slice and make sure that array is modified
	s1[0] = 42
	if a[1] != 42 {
		t.Errorf("array not modified: %v -> %v", a, s1)
	}

	var s2 []int

	if s2 != nil || len(s2) != 0 || cap(s2) != 0 {
		t.Errorf("should be nil slice: %v -> %v -> %v", s2, len(s2), cap(s2))
	}

	// allocate zeroed array of len 5 and return slice of len 2
	s3 := make([]int, 2, 5)

	if s3[0] != 0 || s3[1] != 0 || len(s3) != 2 || cap(s3) != 5 {
		t.Errorf("should be dynamically sized slice: %v -> %v -> %v", s3, len(s3), cap(s3))
	}
}

func TestArrays2D(t *testing.T) {
	// 2D: slice of slices
	field := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	if field[0][0] != 1 {
		t.Errorf("hmmm... incorrect 2d array addressing: %v != 0", field[0][0])
	}

	if field[1][1] != 5 {
		t.Errorf("hmmm... incorrect 2d array addressing: %v != 5", field[0][0])
	}

	if field[1][2] != 6 {
		t.Errorf("hmmm... incorrect 2d array addressing: %v != 6", field[0][0])
	}
}

func TestAppend(t *testing.T) {
	var s []int

	if len(s) != 0 {
		t.Errorf("hmmm... nil slice length is non-zero: %v", len(s))
	}

	s = append(s, 0)
	if len(s) != 1 || s[0] != 0 {
		t.Errorf("hmmm... incorrect slice length and values: %v -> %v", len(s), s)
	}

	s = append(s, 1)
	if len(s) != 2 || s[1] != 1 {
		t.Errorf("hmmm... incorrect slice length and values: %v -> %v", len(s), s)
	}

	s = append(s, 2, 3, 4)
	if len(s) != 5 || s[4] != 4 {
		t.Errorf("hmmm... incorrect slice length and values: %v -> %v", len(s), s)
	}
}

func TestRange(t *testing.T) {
	var pow = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	for i, v := range pow {
		if v != i*i {
			t.Errorf("hmmm... pow[%v] = %v != %v*%v == %v", i, v, i, i, i*i)
		}
	}

	// dynamic allocation
	grid := make([][]int, 5)

	for i := range grid {
		grid[i] = make([]int, 10)
		for j := range grid[i] {
			grid[i][j] = i + j
		}
	}

	if grid[1][2] != 3 {
		t.Errorf("hmmm... incorrect 2d grid value: %v != 3", grid[1][2])
	}
}

type Vertex struct {
	Lat, Long float64
}

func TestMaps(t *testing.T) {
	tm := make(map[string]Vertex)

	_, ok := tm["Bell Labs"]
	if ok != false {
		t.Errorf("hmmm... unexpected map value: %v", tm["test"])
	}

	tm["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	v, ok := tm["Bell Labs"]

	if ok != true {
		t.Errorf("hmmm... missing map value for key %v", "Bell Labls")
	}

	if v.Lat != 40.68433 || v.Long != -74.39967 {
		t.Errorf("hmmm... unexpected map value %v", v)
	}

	delete(tm, "Bell Labs")
	_, ok = tm["Bell Labs"]
	if ok != false {
		t.Errorf("hmmm... unexpected map value: %v", tm["test"])
	}
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, w := range strings.Fields(s) {
		m[w] += 1
	}

	return m
}

func TestWordCount(t *testing.T) {
	m := WordCount("hello world")
	if m["hello"] != 1 || m["world"] != 1 {
		t.Errorf("hmmm... unexpected word count: %v", m)
	}

	m = WordCount("a b c a b a")
	if m["a"] != 3 || m["b"] != 2 || m["c"] != 1 {
		t.Errorf("hmmm... unexpected word count: %v", m)
	}
}
