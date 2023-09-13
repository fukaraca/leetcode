package main

import "fmt"

func main() {
	a := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
		{-1},
	}
	for _, ints := range a {
		fmt.Println(maxSubArray2(ints))
	}

}

func maxSubArray(nums []int) int {
	out := nums[0]
	l := len(nums)
	if l == 0 {
		return out
	}
	for i := 0; i < l; i++ {
		var t int
		for j := i; j < l; j++ {
			t += nums[j]
			if t > out {
				out = t
			} else {
				nums[j] = 0
			}
		}
	}
	return out
}

func maxSubArray2(nums []int) int {
	out := nums[0]
	l := len(nums)
	if l == 0 {
		return out
	}
	var pre int
	for i := 0; i < l; i++ {
		if pre < 0 {
			pre = 0
		}
		pre += nums[i]
		if out < pre {
			out = pre
		}
	}
	return out
}
