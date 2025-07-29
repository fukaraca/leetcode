package main

func main() {
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	if m < 3 || n < 3 {
		return
	}

	type ind struct {
		y, x int
	}

	var dfs func(ik, jk int, v map[ind]bool) bool
	dfs = func(ik, jk int, v map[ind]bool) bool {
		if v[ind{ik, jk}] {
			return false
		}
		if ik == m || jk == n || ik < 0 || jk < 0 {
			return true
		}
		if board[ik][jk] == 'X' {
			return false
		}

		v[ind{ik, jk}] = true
		right := dfs(ik, jk+1, v)
		left := dfs(ik, jk-1, v)
		down := dfs(ik+1, jk, v)
		up := dfs(ik-1, jk, v)
		return right || left || down || up
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				visited := make(map[ind]bool)
				if !dfs(i, j, visited) {
					for idx := range visited {
						board[idx.y][idx.x] = 'X'
					}
				}
			}
		}
	}
}
