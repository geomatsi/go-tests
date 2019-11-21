//
// mix of examples from 'A Tour of Go': section 'Flow control statements'
//

package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrtApprox(x float64) (float64, int) {

	i := 0
	z1 := 1.0
	z2 := 1.0

	for {
		i, z1, z2 = i+1, z1-(z1*z1-x)/(2*z1), z1
		if math.Abs(z1-z2) < 1e-6 {
			return z1, i
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		var v = i
		defer func() {
			switch {
			case v < 3:
				fmt.Println("main")
			case v < 6:
				fmt.Println("main main")
			case v < 9:
				fmt.Println("main main main")
			default:
				fmt.Println("main main main main")
			}
		}()
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Printf("good: %v\n", os)
	case "darwin":
		fmt.Printf("ummh: %v\n", os)
	default:
		fmt.Printf("hmmm: %v\n", os)
	}

	var t1 int

	for i := 0; i < 10; i++ {
		t1 += i
	}

	fmt.Println(t1)

	t2 := 1

	for t2 < 1000 {
		t2 += t2
		if t := t2; t > 100 {
			break
		}
	}

	fmt.Println(t2)

	for i := 0; i < 20; i++ {
		v, n := sqrtApprox(float64(i))
		fmt.Printf("sqrt(%v) = %v, %v iterations\n", i, v, n)
	}
}
