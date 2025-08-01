package main

import "fmt"

func main() {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix2(m, 6))
	fmt.Println(searchMatrix2([][]int{{1}}, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	p1, p2, j := 0, m, n-1
	i := (p1 + p2) / 2
	for p1 < p2 {
		switch {
		case matrix[i][j] < target:
			p1 = i + 1
			i = (p1 + p2) / 2
		case matrix[i][j] > target:
			p2 = i
			i = (p1 + p2) / 2
		default:
			return true
		}
	}
	p1, p2 = 0, n
	j = (p1 + p2) / 2
	for p1 < p2 {
		switch {
		case matrix[i][j] < target:
			p1 = j + 1
			j = (p1 + p2) / 2
		case matrix[i][j] > target:
			p2 = j
			j = (p1 + p2) / 2
		default:
			return true
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	row, col := 0, n-1
	for row < m && col >= 0 {
		switch {
		case matrix[row][col] < target:
			row++
		case matrix[row][col] > target:
			col--
		default:
			return true
		}
	}
	return false
}
