package main

import (
	"fmt"
)

func main() {
	input := [][]int{{1, 3}, {4, 6}, {8, 10}, {15, 18}}
	ins := []int{5, 12}
	fmt.Println(insert2(input, ins))
	input = [][]int{{1, 5}}
	ins = []int{2, 3}
	fmt.Println(insert2(input, ins))

}
func insert(intervals [][]int, newInterval []int) [][]int {
	out := make([][]int, 0)
	var start, end *int
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	for i := 0; i < len(intervals); i++ {
		if start == nil {
			if intervals[i][0] >= newInterval[0] || newInterval[0] <= intervals[i][1] {
				s := mini(newInterval[0], intervals[i][0])
				start = &s
				if newInterval[1] < intervals[i][0] {
					out = append(out, []int{*start, newInterval[1]})
					out = append(out, intervals[i])
					continue
				}
				e := maxi(intervals[i][1], newInterval[1])
				end = &e
				continue
			}
			out = append(out, intervals[i])
			continue
		}
		if end != nil && intervals[i][0] >= *end+1 {
			out = append(out, []int{*start, *end})
			out = append(out, intervals[i])
			end = nil
			continue
		} else if end != nil && intervals[i][1] > *end {
			end = &intervals[i][1]
			continue
		} else if end == nil {
			out = append(out, intervals[i])
		}
	}
	if start == nil {
		out = append(out, newInterval)
	}
	if end != nil {
		out = append(out, []int{*start, *end})
	}
	return out
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	out := make([][]int, 0)
	l := len(intervals)
	var i int

	for i < l && newInterval[0] > intervals[i][1] {
		out = append(out, intervals[i])
		i++
	}

	for i < l && newInterval[1] >= intervals[i][0] {
		//merge
		newInterval[0] = mini(newInterval[0], intervals[i][0])
		newInterval[1] = maxi(newInterval[1], intervals[i][1])
		i++
	}
	out = append(out, newInterval)

	for i < l {
		out = append(out, intervals[i])
		i++
	}
	return out
}

func maxi(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func mini(i, j int) int {
	if i < j {
		return i
	}
	return j
}
