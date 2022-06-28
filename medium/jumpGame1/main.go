package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 0, 0}
	fmt.Println(canJump(nums))

}

func canJump(nums []int) bool {
	far := 0
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if far >= len(nums)-1 {
			return true
		}
		if far == i && nums[i] == 0 {
			return false
		}
	}
	return true

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
