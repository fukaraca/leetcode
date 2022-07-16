package main

import (
	"fmt"
)

func main() {
	x := float64(2)
	n := 21
	fmt.Println("f:", myPow(x, n))

}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	if n < 0 {

		return 1 / (myPow(x, -n))
	}
	if n%2 == 0 {
		ret := myPow(x, n/2)
		return ret * ret
	}

	return myPow(x, n-1) * x
}
