package main

import "fmt"

func main() {
	fmt.Println(categorizeBox(10000, 1000, 1, 13))

}

func categorizeBox(length int, width int, height int, mass int) string {
	v := int64(length) * int64(width) * int64(height)
	b := v >= 1_000_000_000 || length >= 10_000 || width >= 10_000 || height >= 10_000
	h := mass >= 100
	switch {
	case b && h:
		return "Both"
	case b && !h:
		return "Bulky"
	case !b && h:
		return "Heavy"
	default:
		return "Neither"
	}
}
