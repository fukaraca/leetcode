package main

import (
	"fmt"
	"sort"
)

func main() {
	can := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(can, target))

}

/*
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := [][]int{}

	var addNcheck func(c []int, tar, tot int, l []int) bool
	liste := []int{}
	addNcheck = func(c []int, tar, tot int, l []int) bool {

		if tot > tar {
			return false
		}
		if tot == target {
			return true
		}

		for _, candidate := range candidates {
			if len(l) != 0 && candidate < l[len(l)-1] {
				continue
			}
			tot += candidate
			l = append(l, candidate)
			if addNcheck(c, tar, tot, l) {

				ret = append(ret, l)
				tot -= candidate
				return false

			} else {
				l = l[: len(l)-1 : len(l)-1]
				tot -= candidate
			}
		}
		return false
	}
	addNcheck(candidates, target, 0, liste)
	return ret
}
