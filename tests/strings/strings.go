//
// mix of simple strings examples
//

package main

import (
	"fmt"
)

func main() {
	fmt.Println("run tests: go test -cover")
}

func addBinary(a string, b string) string {
	res := ""

	n := len(a)
	if n < len(b) {
		n = len(b)
	}

	x := 0
	y := 0
	z := 0

	for i := 0; i <= n; i++ {

		if i < len(a) {
			switch a[len(a)-i-1] {
			case '0':
				x = 0
			case '1':
				x = 1
			default:
				return ""
			}
		} else {
			x = 0
		}

		if i < len(b) {
			switch b[len(b)-i-1] {
			case '0':
				y = 0
			case '1':
				y = 1
			default:
				return ""
			}
		} else {
			y = 0
		}

		c := x + y + z

		switch c {
		case 0:
			if i < n {
				res = "0" + res
				z = 0
			}
		case 1:
			res = "1" + res
			z = 0
		case 2:
			res = "0" + res
			z = 1
		case 3:
			res = "1" + res
			z = 1
		default:
			return ""
		}
	}

	return res
}
