package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

}

func maxArea(height []int) int {
	max := 0
	size := len(height)
	for i, j := 0, size-1; i <= j; {
		mn := min(height[i], height[j])
		if (j-i)*mn > max {
			max = (j - i) * mn
		} else {
			switch {
			case mn == height[i]:
				i++
			case mn == height[j]:
				j--
			}
		}
	}
	return max
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
