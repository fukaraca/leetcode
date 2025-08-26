package main

func main() {
}

func rotateTheBox(boxGrid [][]byte) [][]byte {

	for i, r := range boxGrid {
		var st []byte
		for j, o := range r {
			if o == '*' {
				for k := 0; k < len(st); k++ {
					boxGrid[i][j-1-k] = st[len(st)-1-k]
				}
				st = []byte{}
			} else if o == '#' {
				st = append(st, o)
				boxGrid[i][j] = '.'
			}
		}
		if len(st) > 0 {
			for k := 0; k < len(st); k++ {
				boxGrid[i][len(r)-1-k] = st[len(st)-1-k]
			}
		}
	}
	out := make([][]byte, len(boxGrid[0]))
	for i := range out {
		out[i] = make([]byte, len(boxGrid))
	}

	for i := len(boxGrid) - 1; i >= 0; i-- {
		for j := 0; j < len(boxGrid[0]); j++ {
			out[j][i] = boxGrid[len(boxGrid)-i-1][j]

		}

	}
	return out
}
