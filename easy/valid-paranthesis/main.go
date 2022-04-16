package main

import "fmt"

func main() {
	a := "({[)"
	b := "[]()[]"
	c := "{[]}"
	fmt.Println(isValid(a), isValid(b), isValid(c))

}

func isValid(s string) bool {
	subopen := false
	open := false
	stack := []int32{}
	inits := map[int32]int32{'(': ')', '[': ']', '{': '}'}
	ends := map[int32]int32{')': '(', ']': '[', '}': '{'}

	if _, ok := ends[int32(s[0])]; ok {
		return false
	}

	for _, c := range s {
		if !subopen && !open {
			if _, ok := inits[c]; ok {
				subopen = true
				open = true
				stack = append(stack, inits[c])

			} else {
				return false
			}
		} else {
			if len(stack) == 1 && c == stack[0] {
				open = false
				subopen = false
				stack = stack[:0]
			} else {
				if c == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					subopen = false
				} else {
					subopen = true
					open = true
					stack = append(stack, inits[c])
				}
			}
		}
	}
	return !open
}
