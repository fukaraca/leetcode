package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
	//aq,ap,ar,as,bq,bp,br,bs
}

func letterCombinations(digits string) []string {
	size := len(digits)
	if size == 0 {
		return []string{}
	}

	buttons := map[int32]string{50: "abc", 51: "def", 52: "ghi", 53: "jkl", 54: "mno", 55: "pqrs", 56: "tuv", 57: "wxyz"}
	retSize := 0
	ret := []string{}
	for _, l := range buttons[int32(digits[0])] {
		retSize++
		ret = append(ret, string(l))
	}

	for _, d := range digits[1:] {
		if d == 55 || d == 57 {
			retSize *= 4
		} else {
			retSize *= 3
		}
		for _, v := range ret {
			for _, l := range buttons[d] {
				ret = append(ret, v+string(l))
			}
		}
	}

	beg := len(ret) - retSize
	return ret[beg:]
}
