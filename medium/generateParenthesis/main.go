package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ret := []string{}
	//backtracking alghorithm
	propagant := func(str string, op, cl int) {}
	propagant = func(str string, op, cl int) {
		//base case
		if op == n && cl == n {
			ret = append(ret, str)
			return
		}
		// create an alternative combination of open paranth
		if op < n {
			propagant(str+"(", op+1, cl)
		}
		// create an alternative combination of close paranth
		if cl < op {
			propagant(str+")", op, cl+1)
		}
	}
	propagant("", 0, 0)

	return ret
}
