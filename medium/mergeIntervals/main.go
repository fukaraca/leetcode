package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(input))
	input = [][]int{{1, 4}, {5, 6}}
	fmt.Println(merge(input))

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := make([][]int, 0)
	curr := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if curr[1]+1 <= intervals[i][0] {
			out = append(out, []int{curr[0], curr[1]})
			curr[0], curr[1] = intervals[i][0], intervals[i][1]
			continue
		}
		if curr[1] < intervals[i][1] {
			curr[1] = intervals[i][1]
			continue
		}
	}
	return append(out, []int{curr[0], curr[1]})
}
