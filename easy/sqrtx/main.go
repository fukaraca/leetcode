package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	mid, left, high := x/2, 0, x
	for left < high {
		m2 := mid * mid
		if m2 == x {
			return mid
		}
		if m2 > x {
			high = mid
		} else {
			left = mid + 1
		}
		mid = (left + high) / 2
	}
	return mid - 1
}
