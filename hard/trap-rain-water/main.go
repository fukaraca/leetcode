package main

import "fmt"

func main() {
	h1 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(h1) == 9)
	h2 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(h2) == 6)
}

func trap(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}
	leftMax := 0
	left := make([]int, l)
	rightMax := 0
	right := make([]int, l)
	sum := 0
	for i, j := 0, l-1; i < l && j >= 0; i, j = i+1, j-1 {
		left[i] = leftMax
		right[j] = rightMax
		if height[i] > leftMax {
			leftMax = height[i]
		}
		if height[j] > rightMax {
			rightMax = height[j]
		}
	}
	for i := 0; i < l; i++ {
		sum += ifPositive(min(left[i], right[i]) - height[i])
	}

	return sum

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func ifPositive(x int) int {
	if x > 0 {
		return x
	}
	return 0
}
