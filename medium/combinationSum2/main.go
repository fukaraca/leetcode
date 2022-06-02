package main

import (
	"fmt"
	"sort"
)

func main() {
	//[[2,5,2,1,2]
	//[10,1,2,7,6,1,5
	//[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
	/*	can := []int{1, 1, 1, 1, 1}
		target := 5*/
	can := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target := 25
	/*	can := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target := 27*/
	/*	can := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8*/
	fmt.Println(combinationSum2(can, target))

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ret := [][]int{}
	var sumAndCheck func(start int, tar int, list []int)

	sumAndCheck = func(start int, tar int, list []int) {
		if tar < 0 {
			return
		}
		if tar == 0 {
			ret = append(ret, append([]int{}, list...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sumAndCheck(i+1, tar-candidates[i], append(list, candidates[i]))
		}
	}
	sumAndCheck(0, target, []int{})
	return ret
}
