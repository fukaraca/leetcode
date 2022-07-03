package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))

}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var ret [][]int

	var dfs func(x, y []int, mp map[int]struct{})

	dfs = func(list, path []int, used map[int]struct{}) {
		if len(path) == len(nums) {
			ret = append(ret, path)
			return
		}
		for _, n := range list {
			if _, ok := used[n]; !ok {
				used[n] = struct{}{}
				path = append(path, n)
				dfs(list, append([]int{}, path...), used)
				path = path[:len(path)-1]
				delete(used, n)
			}
		}
	}

	dfs(nums, []int{}, make(map[int]struct{}))

	return ret
}
