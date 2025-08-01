package main

import "fmt"

func main() {
	a := []int{2, 0, 1}
	sortColors2(a)
	fmt.Println(a)
}

func sortColors(nums []int) {
	ct := make([]int, 3)
	for _, num := range nums {
		ct[num]++
	}
	var j int
	for i := range ct {
		j = i
		break
	}
	for i := 0; i < len(nums); i++ {
		if ct[j] == 0 {
			j++
			i--
			continue
		}
		nums[i] = j
		ct[j]--
	}

}

func sortColors2(nums []int) {
	l, m, h := 0, 0, len(nums)-1

	for m <= h {
		fmt.Println(nums)
		switch nums[m] {
		case 0:
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		case 1:
			m++
		case 2:
			nums[h], nums[m] = nums[m], nums[h]
			h--
		}
	}
}
