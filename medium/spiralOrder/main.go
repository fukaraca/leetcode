package main

import "fmt"

func main() {
	a := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		{{0, 1}, {0, 1}},
		{{0}},
	}
	for _, i2 := range a {
		fmt.Println(spiralOrder(i2))
	}
}

func spiralOrder(matrix [][]int) []int {
	var out []int
	var i, j, li, lj = 0, 0, -1, -1
	var pi, pj = 0, -1
	mi := len(matrix)
	mj := len(matrix[0])
	for i < mi && j < mj && i > li && j > lj {
		out = append(out, matrix[i][j])
		if pi == i && pj < j { // has gone right
			if j+1 != mj {
				pj = j
				j++
				continue
			}
			pj, pi = j, i
			i++
			li++
		} else if pj == j && pi < i { // has gone down
			if i+1 != mi {
				pi = i
				i++
				continue
			}
			pj, pi = j, i
			j--
			mj--
		} else if pi == i && pj > j { // has gone left
			if j-1 != lj {
				pj = j
				j--
				continue
			}
			pj, pi = j, i
			i--
			mi--
		} else { // has gone up
			if i-1 != li {
				pi = i
				i--
				continue
			}
			pj, pi = j, i
			j++
			lj++
		}
	}

	return out
}
