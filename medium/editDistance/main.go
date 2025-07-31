package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	w1 := len(word1)
	w2 := len(word2)
	cache := make(map[[2]int]int)
	var check func(i, j int) int
	check = func(i, j int) int {
		switch {
		case i == w1 && j == w2:
			return 0
		case j == w2:
			return w1 - i
		case i == w1:
			return w2 - j
		}
		if v, ok := cache[[2]int{i, j}]; ok {
			return v
		}

		var minCheck int
		if word1[i] == word2[j] {
			minCheck = check(i+1, j+1)
		} else {
			minCheck = min(check(i, j+1), check(i+1, j), check(i+1, j+1)) + 1
		}
		cache[[2]int{i, j}] = minCheck
		return minCheck
	}
	return check(0, 0)
}
