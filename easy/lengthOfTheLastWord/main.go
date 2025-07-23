package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	var c int
	var started bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if started {
				return c
			}
			continue
		}
		started = true
		c++
	}
	return c
}
