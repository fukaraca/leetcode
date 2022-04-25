package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(a, val))

}

func removeElement(nums []int, val int) int {
	ind := 0

	for _, num := range nums {
		if num != val {
			nums[ind] = num
			ind++
		}
	}
	return ind
}
