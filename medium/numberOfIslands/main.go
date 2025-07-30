package main

func main() {
}

func numIslands(grid [][]byte) int {
	var ct int
	m := len(grid)
	n := len(grid[0])
	type ind struct {
		y, x int
	}
	var dfs func(y, x int)
	dfs = func(y, x int) {
		if y == m || y < 0 || x == n || x < 0 || grid[y][x] == '0' {
			return
		}
		grid[y][x] = '0'
		dfs(y, x+1)
		dfs(y+1, x)
		dfs(y, x-1)
		dfs(y-1, x)
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ct++
			}
		}
	}
	return ct
}
