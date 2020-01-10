//
// mix of examples and exercises from 'A Tour of Go': section 'Methods'
//

package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"strings"
	"testing"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func TestVertexAbs(t *testing.T) {
	var v Vertex = Vertex{3.0, 4.0}
	var p *Vertex = &v

	assert.Equal(t, 5.0, v.Abs())
	assert.Equal(t, 5.0, Abs(v))
	assert.Equal(t, 5.0, p.Abs())
	assert.Equal(t, 5.0, Abs(*p))
}

func (v *Vertex) Scale(s float64) {
	v.x *= s
	v.y *= s
}

func Scale(v *Vertex, s float64) {
	v.x *= s
	v.y *= s
}

func TestVertexScale(t *testing.T) {
	var u Vertex = Vertex{1, 2}
	var p *Vertex = &u

	u.Scale(2)
	assert.Equal(t, Vertex{2, 4}, u)

	Scale(&u, 2)
	assert.Equal(t, Vertex{4, 8}, u)

	p.Scale(2)
	assert.Equal(t, Vertex{8, 16}, u)

	Scale(p, 2)
	assert.Equal(t, Vertex{16, 32}, u)
}

type I interface {
	M() int
}

type V struct {
	S string
}

func (v V) M() int {
	return len(v.S)
}

type W struct {
	S string
}

func (w *W) M() int {
	if w == nil {
		return 0
	}

	return len(w.S)
}

func TestImplicitInterface(t *testing.T) {
	var s string = "hello"
	var v V = V{s}

	assert.Equal(t, len(s), v.M())

	var w W = W{s}

	assert.Equal(t, len(s), w.M())

	var i I = V{s}

	assert.Equal(t, len(s), i.M())

	var p I = v

	assert.Equal(t, len(s), p.M())

	var e *W
	var q I = e

	assert.Equal(t, 0, q.M())

	q = &w
	assert.Equal(t, len(s), q.M())
}

func TestTypeAssertions(t *testing.T) {
	var s string = "hello"
	var i I = V{s}
	// nil specific to *W
	var e *W

	v, ok := i.(V)

	assert.Equal(t, true, ok)
	assert.Equal(t, V{s}, v)

	w, err := i.(*W)

	assert.Equal(t, false, err)
	assert.Equal(t, e, w)
}

func typeSwitch(e interface{}) (res int) {
	switch e.(type) {
	case int:
		res = 0
	case string:
		res = 1
	case bool:
		res = 2
	default:
		res = 3
	}

	return
}

func TestTypeSwitch(t *testing.T) {
	var v V

	assert.Equal(t, 0, typeSwitch(1))
	assert.Equal(t, 1, typeSwitch("hello"))
	assert.Equal(t, 2, typeSwitch(false))

	assert.Equal(t, 3, typeSwitch(2.0))
	assert.Equal(t, 3, typeSwitch(v))
}

// Exercise:
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

func TestStringer(t *testing.T) {
	var loopback IPAddr = IPAddr{127, 0, 0, 1}
	var google IPAddr = IPAddr{8, 8, 8, 8}

	assert.Equal(t, "127.0.0.1", fmt.Sprint(loopback))
	assert.Equal(t, "8.8.8.8", google.String())
}

// Exercise:
// Custom error type

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	var err string

	if float64(e) < 0 {
		err = fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
	}

	return err
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z1 = 1.0
	var z2 = 1.0

	for {
		z1, z2 = z1-(z1*z1-x)/(2*z1), z1
		if math.Abs(z1-z2) < 1e-6 {
			return z1, nil
		}
	}
}

func TestCustomError(t *testing.T) {
	v1, e1 := Sqrt(2)

	assert.Equal(t, math.Sqrt(2), v1)
	assert.Equal(t, nil, e1)

	v2, e2 := Sqrt(-2)

	assert.Equal(t, 0.0, v2)
	assert.Equal(t, ErrNegativeSqrt(-2), e2)
	assert.Equal(t, "cannot sqrt negative number: -2", e2.Error())
}

// Exercise:
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'

type StreamReader struct{}

func (r StreamReader) Read(b []byte) (n int, err error) {

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	return len(b), nil
}

func TestStreamReader(t *testing.T) {
	var r StreamReader
	var b []byte

	r.Read(b)
	assert.Equal(t, "", string(b))

	b = make([]byte, 1)
	r.Read(b)
	assert.Equal(t, "A", string(b))

	b = make([]byte, 2)
	r.Read(b)
	assert.Equal(t, "AA", string(b))

	b = make([]byte, 10)
	r.Read(b)
	assert.Equal(t, "AAAAAAAAAA", string(b))
}

// Exercise: rot13Reader
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
// modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

type rot13Reader struct {
	r io.Reader
}

func caesar(v byte, n byte) byte {
	switch {
	case v >= 'a' && v <= 'z':
		return 'a' + (v+n-'a')%('z'-'a'+1)
	case v >= 'A' && v <= 'Z':
		return 'A' + (v+n-'A')%('Z'-'A'+1)
	default:
		return v
	}
}

func (c *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = c.r.Read(b)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		b[i] = caesar(b[i], 13)
	}

	return
}

func TestRot13Reader(t *testing.T) {
	dict := map[string]string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz": "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm",
		"Lbh penpxrq gur pbqr!":                                "You cracked the code!",
		"Jul qvq gur puvpxra pebff gur ebnq?":                  "Why did the chicken cross the road?",
		"Gb trg gb gur bgure fvqr!":                            "To get to the other side!",
		"HELLO":                                                "URYYB",
		"hello":                                                "uryyb",
		"!@#$%^&*()_":                                          "!@#$%^&*()_",
	}

	for enc, dec := range dict {
		s := strings.NewReader(enc)
		r := rot13Reader{s}
		b := make([]byte, len(enc))

		r.Read(b)
		assert.Equal(t, dec, string(b))
	}
}
