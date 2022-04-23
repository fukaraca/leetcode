package main

import "fmt"

func main() {
	a := []int{1, 1, 2}
	fmt.Println(removeDuplicates(a))
}

func removeDuplicates(nums []int) int {
	var temp = -100
	var counter int
	for _, num := range nums {
		if num != temp {
			nums[counter] = num
			counter++
			temp = num

		}
	}
	return counter
}
