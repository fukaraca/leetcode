package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, -1, 0, 2, -2, 0}
	fmt.Println(fourSum(a, 0))

}

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return [][]int{}
	}
	ret := [][]int{}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < size-3; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		temp := threeSum(nums[i+1:], nums[i], target)
		for _, ints := range temp {
			ret = append(ret, ints)
		}
	}

	return ret
}

func threeSum(nums []int, one int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var ret [][]int
	//sort.Ints(nums)
	for i, num := range nums {
		//to fix only unique number
		if i != 0 && nums[i-1] == num {
			continue
		}
		//we fixed i, so we will make closer j and k conditionally
		for j, k := i+1, len(nums)-1; j < k; {
			n := num + nums[j] + nums[k] + one

			if n == target {
				ret = append(ret, []int{num, nums[j], nums[k], one})
				l := j
				for l < k && nums[l] == nums[j] {
					l++
				}
				j = l
			} else if n > target { //this means distribution is weighted to positive so we need to move left
				k--
			} else {
				//move right
				j++
			}
		}
	}
	return ret
}
