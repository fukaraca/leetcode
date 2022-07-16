package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}, {13, 14, 15, 16}}
	rotate(mat)
	fmt.Println(mat)
}

func rotate(matrix [][]int) {
	l := len(matrix[0])
	if l == 1 {
		return
	}
	//first take transpose later symmetry
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-1-j] = matrix[i][l-1-j], matrix[i][j]
		}
	}
}
