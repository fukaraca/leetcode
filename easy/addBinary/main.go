package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
	fmt.Println('a', '0', ' ')
}

func addBinary(a string, b string) string {
	ia, ib := len(a)-1, len(b)-1
	var out []bool
	var sum = uint8(48)
	for ia >= 0 || ib >= 0 {
		if ia >= 0 {
			sum += a[ia]
			ia--
		} else {
			sum += 48
		}
		if ib >= 0 {
			sum += b[ib]
			ib--
		} else {
			sum += 48
		}
		switch sum {
		case 147:
			out = append(out, true)
			sum = 49
		case 146:
			out = append(out, false)
			sum = 49
		case 145:
			out = append(out, true)
			sum = 48
		default:
			out = append(out, false)
			sum = 48
		}
	}
	if sum > 48 {
		out = append(out, true)
	}
	var sb strings.Builder
	for i := len(out) - 1; i >= 0; i-- {
		if out[i] {
			sb.WriteRune(49)
		} else {
			sb.WriteRune(48)
		}
	}
	return sb.String()
}
