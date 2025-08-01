package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(a))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var out [][]int
	n := len(nums)
	var dfs func(i int, arr []int)
	dfs = func(i int, arr []int) {
		t := make([]int, len(arr))
		copy(t, arr)
		out = append(out, t)
		pre := -11
		for ; i < n; i++ {
			if nums[i] == pre {
				continue
			}
			pre = nums[i]
			arr = append(arr, pre)
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0, []int{})
	return out
}
