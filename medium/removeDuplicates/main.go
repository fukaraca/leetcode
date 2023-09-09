package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3})) // 5
}

func removeDuplicates(nums []int) int {
	var temp = nums[0]
	var counter int
	var down = 2
	for i := range nums {
		if nums[i] != temp {
			if down < 2 {
				down = 2
			}
			nums[counter] = nums[i]
			counter++
			down--
			temp = nums[i]
		} else if down == 0 {
		} else {
			nums[counter] = nums[i]
			counter++
			down--
		}

	}
	return counter
}

func removeDuplicates2(nums []int) int {
	var counter = 2
	l := len(nums)
	if l < 2 {
		return l
	}
	for i := 2; i < l; i++ {
		if nums[i] != nums[counter-2] {
			nums[counter] = nums[i]
			counter++
		}
	}
	fmt.Println(nums)
	return counter
}
