package main

import "math"

func main() {
}

func minimumOperationsToWriteY(grid [][]int) int {
	r := make(map[int]int)
	y := make(map[int]int)
	l := len(grid)
	cr, cy := 0, 0

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if (i == j && i <= l/2) || (j > i && j == l-1-i) || (i >= l/2 && j == l/2) {
				y[grid[i][j]]++
				cy++
			} else {
				r[grid[i][j]]++
				cr++
			}

		}
	}
	var cand = math.MaxInt
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			cand = mini(cy-y[i]+cr-r[j], cand)
		}
	}
	return cand
}

func mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}
