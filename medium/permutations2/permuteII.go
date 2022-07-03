package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 3, 0, 3}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var ret [][]int
	used := make(map[int]int)
	var dfs func(y []int, mp map[int]int)
	for _, num := range nums {
		used[num]++
	}
	dfs = func(path []int, used map[int]int) {
		if len(path) == len(nums) {
			ret = append(ret, path)
			return
		}
		for n, v := range used {
			if v > 0 {
				path = append(path, n)
				used[n]--
				dfs(append([]int{}, path...), used)
				used[n]++
				path = path[:len(path)-1]
			}
		}
	}
	dfs([]int{}, used)
	return ret
}
