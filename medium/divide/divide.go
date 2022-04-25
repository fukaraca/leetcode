package main

import (
	"fmt"
	"math"
)

func main() {
	divident := 7
	divisor := 0
	fmt.Println(divide(divident, divisor))

}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return int(math.Pow(2, 31)) - 1
	}
	if (dividend / divisor) > int(math.Pow(2, 31))-1 {
		return int(math.Pow(2, 31)) - 1
	} else if (dividend / divisor) < (-1)*(int(math.Pow(2, 31))) {
		return (-1) * (int(math.Pow(2, 31)))
	}
	return dividend / divisor
}
