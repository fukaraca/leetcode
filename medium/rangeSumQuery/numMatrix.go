package main

type NumMatrix struct {
	matrix [][]int
}

//create presummed matrix
func Constructor(matrix [][]int) NumMatrix {
	lm := len(matrix) + 1
	lm0 := len(matrix[0]) + 1
	nM := make([][]int, lm)
	for i := 0; i < lm; i++ {
		nM[i] = make([]int, lm0)
	}

	for i := 1; i < lm; i++ {
		for j := 1; j < lm0; j++ {
			nM[i][j] = matrix[i-1][j-1] + nM[i][j-1] + nM[i-1][j] - nM[i-1][j-1]
		}

	}
	return NumMatrix{matrix: nM}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.matrix[row2+1][col2+1] - this.matrix[row2+1][col1] - this.matrix[row1][col2+1] + this.matrix[row1][col1]

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
