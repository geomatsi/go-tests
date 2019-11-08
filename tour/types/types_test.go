//
// mix of examples from 'A Tour of Go': section 'Flow control statements'
//

package main

import (
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
