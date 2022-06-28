package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	nums2 := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
	fmt.Println(jump(nums2))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	j := 0
	end, far := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, nums[i]+i) //check if you can jump from here to farthest

		if i == end { //update end to check further jump-point candidates
			j++
			end = far
		}
	}
	return j
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
