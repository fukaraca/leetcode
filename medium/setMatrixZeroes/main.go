package main

func main() {
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	cachj := make(map[int]bool)
	cachi := make(map[int]bool)

	for i := 0; i < m; i++ {
		var clean bool
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				clean = true
				cachi[i], cachj[j] = true, true
			}
		}
		if clean {
			matrix[i] = make([]int, n)
		}
	}

	for i := 0; i < m; i++ {
		if cachi[i] {
			continue
		}
		for j := range cachj {
			matrix[i][j] = 0
		}
	}
}
