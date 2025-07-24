package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(grid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			var vi, vj int
			if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else {
				vi = obstacleGrid[i-1][j]
			}

			if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else {
				vj = obstacleGrid[i][j-1]
			}

			obstacleGrid[i][j] = vi + vj
		}
	}
	return obstacleGrid[m-1][n-1]
}
