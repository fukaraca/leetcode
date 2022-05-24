package main

import "fmt"

/* Given an array of integers nums sorted in non-decreasing order, \
find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.*/

func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums1 := []int{2, 2}
	target := 2
	fmt.Println(searchRange(nums1, target))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	l := 0
	h := len(nums) - 1
	for h > l {
		ind := (h + l) / 2
		if nums[ind] < target {
			l = ind + 1
		} else if nums[ind] > target {
			h = ind - 1
		} else {
			if ind != 0 && nums[ind-1] == target {
				h = ind - 1
				continue
			}
			l = ind
			break
		}
	}

	if nums[l] != target {
		return []int{-1, -1}
	}
	for i := l; i < len(nums) && target == nums[i]; i++ {
		h = i
	}
	return []int{l, h}
}
