package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(5))

}

func fizzBuzz(n int) []string {
	out := make([]string, n, n)
	for i := 0; i < n; i++ {
		switch {
		case (i+1)%15 == 0:
			out[i] = "FizzBuzz"
		case (i+1)%5 == 0:
			out[i] = "Buzz"
		case (i+1)%3 == 0:
			out[i] = "Fizz"
		default:
			out[i] = strconv.Itoa((i + 1))
		}
	}
	return out
}
