package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 0}
	b := []int{1}
	c := []int{7, 8, 9, 11, 12}
	d := []int{3, 4, -1, 1}
	f := [][]int{a, b, c, d}
	for _, ints := range f {
		fmt.Println(firstMissingPositive(ints))
	}

}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	//make all negatives zero
	for i := 0; i < l; i++ {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}
	//save all existed values to indices as negative, save all zeros to beyond length
	for i := 0; i < l; i++ {
		if val := abs(nums[i]); val != 0 && val <= l {
			if nums[val-1] > 0 {
				nums[val-1] = -nums[val-1]
			} else if nums[val-1] == 0 {

				nums[val-1] = -(l + 1)
			}
		}

	}
	ind := 1
	for ; ind <= l; ind++ {
		if nums[ind-1] >= 0 {
			return ind
		}
	}

	return ind
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//with extra space
func firstMissingPositive1(nums []int) int {

	intMap := make(map[int]struct{})
	for i := 1; i <= len(nums); i++ {
		intMap[i] = struct{}{}
	}
	for _, num := range nums {
		if num > 0 && num <= len(nums) {
			if _, ok := intMap[num]; ok {
				delete(intMap, num)
			}
		}
	}
	smallest := len(nums) + 1
	for k, _ := range intMap {
		if k < smallest {
			smallest = k
		}
	}

	return smallest
}
