package main

import (
	"fmt"
)

//not solved yet
func main() {
	nums := []int{1, 1, 3, 4}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var ret [][]int

	var dfs func(x, y []int, mp map[int]int)
	used := make(map[int]int)
	for _, num := range nums {
		used[num]++
	}
	dfs = func(list, path []int, used map[int]int) {
		if len(path) == len(nums) {
			ret = append(ret, path)
			return
		}
		for n, counter := range used {
			for i := counter; i >= 0; i-- {
				used[n]++
				path = append(path, n)
				dfs(list, append([]int{}, path...), used)
				path = path[:len(path)-1]
				used[n]--
			}
		}
	}

	dfs(nums, []int{}, used)

	return ret
}
