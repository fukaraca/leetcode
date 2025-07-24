package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	mt := make([][]int, 0)
	for i := 0; i < m; i++ {
		mt = append(mt, make([]int, n))
		for j := 0; j < n; j++ {
			var vi, vj int
			if i == 0 || j == 0 {
				vi = 1
			} else {
				vi = mt[i-1][j]
				vj = mt[i][j-1]
			}
			mt[i][j] = vi + vj
		}
	}
	return mt[m-1][n-1]
}
