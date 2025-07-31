package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert2("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {

	size := len(s)
	if size == 1 {
		return s
	} else if size == 2 && numRows == 2 {
		return fmt.Sprintf("%s\n%s", string(s[0]), string(s[1]))
	} else if size == 2 && numRows == 1 {
		return s
	} else if numRows == 1 {
		return s
	}
	period := numRows*2 - 2
	column := (size / period) * period
	residue := size - (size/period)*period
	if residue > numRows {
		column += residue - numRows + 1
	} else if residue > 0 && residue <= numRows {
		column += 1
	}

	arr := make([][]int32, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int32, column)
	}
	r, c, q := 0, 0, numRows*2-2
	sb := strings.Builder{}
	for i, l := range s {
		fmt.Println(string(l), column, residue, r, c)
		arr[r][c] = l
		switch {
		case i%q > 0 && i%q < numRows-1:
			r++
		case i%q == numRows-1:
			r--
			c++
		case i%q > numRows-1:
			r--
			c++
		case i%q == 0:
			r++
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < column; j++ {
			if arr[i][j] > 0 {
				sb.WriteString(string(arr[i][j]))
			} else {
				//sb.WriteString(" ")
			}
		}
		//sb.WriteString("\n")
	}

	return sb.String()
}

func convert2(s string, numRows int) string {
	lines := make([]string, numRows)
	var backward bool
	var lidx int

	for i := range s {
		lines[lidx] += string(s[i])

		switch {
		case backward:
			if lidx > 0 {
				lidx--
				continue
			}
			lidx++
			backward = false
		case lidx < numRows-1:
			lidx++
		default:
			lidx--
			backward = true
		}

	}
	var sb strings.Builder
	for i := range lines {
		sb.WriteString(lines[i])
	}
	return sb.String()
}
