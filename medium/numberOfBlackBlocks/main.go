package main

func main() {
}

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	out := []int64{0, 0, 0, 0, 0}
	blockSize := (m - 1) * (n - 1)
	type a struct {
		x, y int
	}
	mp := make(map[a]int)

	for _, v := range coordinates {
		x := v[0]
		y := v[1]

		for i := x - 1; i < m-1 && i < x+1; i++ {
			if i < 0 {
				i++
			}
			for j := y - 1; j < n-1 && j < y+1; j++ {
				if j < 0 {
					j++
				}
				mp[a{i, j}]++
			}
		}
	}

	for _, v := range mp {
		out[v]++
	}
	out[0] = int64(blockSize - len(mp))
	return out

}
