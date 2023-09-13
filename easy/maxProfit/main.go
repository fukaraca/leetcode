package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
}

func maxProfit(prices []int) int {
	var max int
	l := prices[0]
	for i := 0; i < len(prices); i++ {
		if l > prices[i] {
			l = prices[i]
		} else if max < prices[i]-l {
			max = prices[i] - l
		}
	}
	return max
}
