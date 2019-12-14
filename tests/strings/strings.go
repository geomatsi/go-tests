//
// mix of simple strings examples
//

package main

import (
	"fmt"
	"strings"
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

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)

	if nlen == 0 {
		return 0
	}

	if nlen > hlen {
		return -1
	}

	for i := 0; i <= (hlen - nlen); i++ {
		sub := haystack[i:(nlen + i)]
		if sub == needle {
			return i
		}
	}

	return -1
}

func longestCommonPrefix(strs []string) string {
	res := ""

	if len(strs) == 0 {
		return res
	}

	xs := strs[0]
	for n := 1; n <= len(xs); n++ {
		prefix := xs[0:n]
		for _, s := range strs {
			if len(s) < n || xs[n-1:n] != s[n-1:n] {
				return res
			}
		}

		res = prefix
	}

	return res
}

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}

	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i, j = i+1, j-1
	}

	return
}

func reverseWordsOrder(s string) string {
	var v []byte

	i := len(s) - 1
	j := len(s) - 1

	for {
		for {
			if i < 0 || s[i] != ' ' {
				break
			}

			i--
		}

		if i < 0 {
			break
		}

		if len(v) > 0 {
			v = append(v, ' ')
		}

		j = i

		for {
			if j < 0 || s[j] == ' ' {
				break

			}

			j--
		}

		for k := j + 1; k < i+1; k++ {
			v = append(v, s[k])

		}

		i = j
	}

	return string(v)
}

func reverseWords(s string) string {
	var v []byte

	i := 0
	j := 0

	for {
		for {
			if i >= len(s) || s[i] != ' ' {
				break
			}

			i++
		}

		if i >= len(s) {
			break
		}

		if len(v) > 0 {
			v = append(v, ' ')
		}

		j = i

		for {
			if j >= len(s) || s[j] == ' ' {
				break

			}

			j++
		}

		for k := j - 1; k >= i; k-- {
			v = append(v, s[k])
		}

		i = j
	}

	return string(v)
}

func reverseAsRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseWordsStd(s string) string {
	strs := strings.Fields(s)
	for p, e := range strs {
		strs[p] = reverseAsRunes(e)
	}

	return strings.Join(strs, " ")
}
