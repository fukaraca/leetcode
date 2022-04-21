package main

import (
	"fmt"
	"sort"
)

// not solved yet
func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(threeSumMulti(arr, 8))
	fmt.Println(combination(3))
}

type helper struct {
	val, ind int
}

func threeSumMulti(arr []int, target int) int {
	var ret int
	size := len(arr)
	var narr []helper

	for i, val := range arr {
		narr = append(narr, helper{val: val, ind: i})
	}
	sort.Slice(narr, func(i, j int) bool {
		if narr[i].ind < narr[j].ind {
			return true
		}
		if narr[i].ind > narr[j].ind {
			return false
		}
		return narr[i].val < narr[j].val
	})
	comb1 := 1
	comb2 := 1
	for i := 0; i < size-2; i++ {
		if narr[i].val == narr[i+1].val && i != size-3 {
			comb1++
			continue
		} else if i == size-3 && narr[i].val == narr[i+1].val && size != 3 {
			comb1++
		}
		for j, k := i+1, size-1; j < k; {
			n := narr[i].val + narr[j].val + narr[k].val
			if n == target && narr[i].ind < narr[j].ind && narr[j].ind < narr[k].ind {

				m := k
				for j < m && narr[m].val == narr[k].val && narr[m].ind > narr[j].ind {
					if m < k {
						comb2++
					}
					m--

				}
				if comb1 == comb2 && comb1 == 1 {
					ret++
				} else {
					ret += combination(comb1 + comb2)
				}
				k = m

				comb1, comb2 = 1, 1
			} else if k-1 == j {
				j++
				k = size - 1
			} else {
				k--

			}

		}
		comb1 = 1
		comb2 = 1

	}
	return ret%10 ^ 9
}
func factorialCal(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	return number * factorialCal(number-1)
}
func combination(x int) int {
	ret := factorialCal(x) / (factorialCal(x-3) * 6)
	return ret
}
