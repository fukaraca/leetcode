package main

import "fmt"

func main() {
	a := []string{"(())(", "(()", ")()())", "((()(())))", "", "()()", "()(()"}
	for _, s := range a {
		fmt.Println(longestValidParentheses(s))
	}
}

var longest int

func longestValidParentheses(s string) int {
	longest = 0
	l := len(s)
	if l == 1 || (l == 2 && s != "()") {
		return 0
	}
	var i int
	for i < l {
		if s[i] == '(' {
			i += lookValid(s[i:]) + 1
			continue
		}
		i++
	}
	return longest
}

func lookValid(s string) int {
	var stack []uint8
	var ct int
	defer func() {
		if longest < ct*2 {
			longest = ct * 2
		}
	}()
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return i
		}
		if stack[len(stack)-1] == '(' {
			ct++
			stack = stack[:len(stack)-1]
			if longest < ct*2 && i+1 == ct*2 {
				longest = ct * 2
			}
			continue
		}

	}
	if len(stack) > 0 {
		ct = 0
	}
	return 0
}
