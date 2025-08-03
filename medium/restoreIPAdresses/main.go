package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))

}

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	var out []string
	var fn func(i, bl int, can string)
	fn = func(i, bl int, can string) {
		if bl == 0 {
			if i < len(s) {
				return
			}
			out = append(out, can)
			return
		}
		bl--
		if i < len(s) {
			fn(i+1, bl, can+"."+s[i:i+1])
		}
		if i < len(s)-1 && isValid(s[i:i+2]) {
			fn(i+2, bl, can+"."+s[i:i+2])
		}
		if i < len(s)-2 && isValid(s[i:i+3]) {
			fn(i+3, bl, can+"."+s[i:i+3])
		}
	}

	for i, s2 := range []string{s[:1], s[:2], s[:3]} {
		if isValid(s2) {
			fn(i+1, 3, s2)
		}
	}
	return out
}
func isValid(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	var out int
	for i := len(s) - 1; i >= 0; i-- {
		out += int(s[i]-'0') * int(math.Pow10(len(s)-i-1))
	}
	return out > -1 && out < 256
}
