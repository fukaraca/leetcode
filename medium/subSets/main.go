package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var out [][]int
	l := len(nums)

	var dfs func(i int, arr []int)
	dfs = func(i int, arr []int) {
		t := make([]int, len(arr))
		copy(t, arr)
		out = append(out, t)
		for ; i < l; i++ {
			arr = append(arr, nums[i])
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	dfs(0, []int{})
	return out
}
