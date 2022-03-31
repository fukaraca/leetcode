package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var ret [][]int
	sort.Ints(nums)
	for i, num := range nums {
		//to fix only unique number
		if i != 0 && nums[i-1] == num {
			continue
		}
		//we fixed i, so we will make closer j and k conditionally
		for j, k := i+1, len(nums)-1; j < k; {
			n := num + nums[j] + nums[k]
			if n == 0 {
				ret = append(ret, []int{num, nums[j], nums[k]})
				l := j
				for l < k && nums[l] == nums[j] {
					l++
				}
				j = l
			} else if n > 0 { //this means distribution is weighted to positive so we need to move left
				k--
			} else {
				//move right
				j++
			}

		}

	}

	return ret
}
