package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7}
	targ := 2
	fmt.Println(searchInsert(nums, targ))

}

func searchInsert(nums []int, target int) int {
	ln := len(nums)

	if target > nums[ln-1] {
		return ln
	}
	if target < nums[0] {
		return 0
	}

	l := 0
	h := ln - 1

	for l <= h {
		ind := (h + l) / 2
		if nums[ind] > target {
			h = ind - 1
			continue
		}
		if nums[ind] < target {
			l = ind + 1
			continue
		}
		if nums[ind] == target {
			return ind
		}
	}

	return l
}
