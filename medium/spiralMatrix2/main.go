package main

import "fmt"

func main() {
	for _, is := range generateMatrix(20) {
		fmt.Println(is)

	}

}

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	if n == 1 {
		return [][]int{{1}}
	}
	for i := range out {
		out[i] = make([]int, n)
	}
	v, i, j := 1, 0, 0
	dir := right
	m := n * n
	for v <= m {
		out[i][j] = v
		v++
		if v <= m {
			i, j, dir = bump(i, j, dir, n, out)
		}

	}

	return out
}

const (
	right = iota
	left
	down
	up
)

func bump(i, j, dir, n int, out [][]int) (int, int, int) {
	switch dir {
	case right:
		if j+1 == n || out[i][j+1] > 0 {
			return bump(i, j, down, n, out)
		}
		return i, j + 1, right
	case down:
		if i+1 == n || out[i+1][j] > 0 {
			return bump(i, j, left, n, out)
		}
		return i + 1, j, down
	case left:
		if j-1 == -1 || out[i][j-1] > 0 {
			return bump(i, j, up, n, out)
		}
		return i, j - 1, left
	default:
		if i-1 == -1 || out[i-1][j] > 0 {
			return bump(i, j, right, n, out)
		}
		return i - 1, j, up
	}
}
