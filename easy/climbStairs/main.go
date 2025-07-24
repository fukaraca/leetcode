package main

import "fmt"

func main() {
	fmt.Println(climbStairs(6))
}

func climbStairs(n int) int {
	sum, pre := 1, 0
	for i := 1; i <= n; i++ {
		t := sum
		sum += pre
		pre = t
	}
	return sum
}
