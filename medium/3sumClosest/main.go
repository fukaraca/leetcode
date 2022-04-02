package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(a, target))

}

func threeSumClosest(nums []int, target int) int {

	var ret = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i, num := range nums {
		//to fix only unique number
		if i != 0 && nums[i-1] == num {
			continue
		}
		//we fixed i, so we will make closer j and k conditionally
		for j, k := i+1, len(nums)-1; j < k; {
			n := num + nums[j] + nums[k]
			if n == target {
				return n
			} else if n > target { //this means distribution is weighted to positive so we need to move left
				k--
			} else {
				//move right
				j++
			}
			ret = check(n, target, ret)

		}

	}

	return ret
}

func check(n, target, ret int) int {
	if abs(ret-target) > abs(n-target) {
		return n
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
