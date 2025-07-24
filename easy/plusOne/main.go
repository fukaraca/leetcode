package main

import "fmt"

func main() {
	d := []int{4, 3, 2, 1}
	fmt.Println(plusOne(d))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
