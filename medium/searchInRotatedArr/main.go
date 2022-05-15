package main

import "fmt"

func main() {
	//  0  2 11 12 13 14 15 16 17 18
	// 18  0  2 11 12 13 14 15 16 17
	// 11 12 13 14 15 16 17 18  0  2
	a := []int{4, 5, 6, 7, 0, 1, 2}
	targ := 3
	fmt.Println(search(a, targ))
}

func search(nums []int, target int) int {
	ln := len(nums)
	if ln == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	var h = ln - 1
	var l = 0

	for h > l {
		ind := (h + l) / 2
		if nums[ind] > nums[h] {
			l = ind + 1
		} else {
			h = ind
		}
	}
	piv := l
	h = ln - 1
	l = 0
	for l <= h {
		mid := (h + l) / 2
		realMid := (mid + piv) % ln
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return -1
}
