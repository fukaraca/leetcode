package flipAndInvertImage

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	if len(image) == 1 {
		image[0][0] = inverse(image[0][0])
		return image
	}
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; k > j; j, k = j+1, k-1 {
			image[i][j], image[i][k] = inverse(image[i][k]), inverse(image[i][j])

		}
		if n%2 != 0 {
			image[i][n/2] = inverse(image[i][n/2])
		}

	}
	return image

}
func inverse(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
