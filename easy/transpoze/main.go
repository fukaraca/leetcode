package main

import "fmt"

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(a))
}

func transpose(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		ret[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ret[j][i] = matrix[i][j]
		}
	}
	return ret
}
