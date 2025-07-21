package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-50050500))
}

func reverse(x int) int {
	var out int
	p := func() int {
		if x >= 0 {
			return 1
		}
		return -1
	}()

	for x != 0 {
		r := x % 10
		x /= 10
		if (p > 0 && (math.MaxInt32-r)/10 >= out) || (p < 0 && (math.MinInt32-r)/10 <= out) {
			out = out*10 + r
		} else {
			return 0
		}
	}
	return out
}
