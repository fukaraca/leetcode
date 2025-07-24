package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println(minPathSum(grid))
	fmt.Println(minPathSum2(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mps := make([][]int, 0)

	for i := 0; i < m; i++ {
		mps = append(mps, make([]int, n))
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				mps[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				mps[i][j] = mps[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				mps[i][j] = mps[i-1][j] + grid[i][j]
				continue
			}
			mps[i][j] = mini(mps[i-1][j], mps[i][j-1]) + grid[i][j]

		}
	}
	return mps[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += mini(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}
