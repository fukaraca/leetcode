package main

import (
	"fmt"
)

func main() {
	s := []string{
		"ABCE",
		"SFCS",
		"ADEE"}
	var b [][]byte
	for _, s2 := range s {
		b = append(b, []byte(s2))
	}
	w := "ABCB"
	fmt.Println(exist(b, w))
	fmt.Println(exist([][]byte{[]byte("a")}, "a"))
	fmt.Println(exist([][]byte{[]byte("ab"), []byte("cd")}, "acdb"))
}

func exist(board [][]byte, word string) bool {
	// iterate over b , start if matched, return if found
	lr := len(board)
	lc := len(board[0])
	lw := len(word)
	type coord struct {
		i, j int
	}
	visited := make([][]bool, lr)
	for i := range visited {
		visited[i] = make([]bool, lc)
	}
	var fn func(int, int, int) bool
	fn = func(i, j, w int) bool {
		if w == lw {
			return true
		}
		if i < 0 || i == lr || j < 0 || j == lc || visited[i][j] || board[i][j] != word[w] {
			return false
		}
		visited[i][j] = true
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range directions {
			if fn(i+dir[0], j+dir[1], w+1) {
				return true
			}
		}
		visited[i][j] = false
		return false
	}

	for i := 0; i < lr; i++ {
		for j := 0; j < lc; j++ {
			if board[i][j] == word[0] {
				if fn(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
