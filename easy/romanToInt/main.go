package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(romanToInt("IV"))
}

func romanToInt(s string) int {
	size := len(s)
	ret := 0
	hrt := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	if size == 1 {
		return hrt[rune(s[0])]
	}

	for i := size - 1; i >= 1; i-- {

		if hrt[rune(s[i])] > hrt[rune(s[i-1])] && i != 1 {
			ret -= hrt[rune(s[i-1])]
			ret += hrt[rune(s[i])]
			i--
		} else if hrt[rune(s[i])] > hrt[rune(s[i-1])] && i == 1 {
			ret -= hrt[rune(s[i-1])]
			ret += hrt[rune(s[i])]
			return ret
		} else {
			ret += hrt[rune(s[i])]
		}
	}
	ret += hrt[rune(s[0])]
	return ret
}
