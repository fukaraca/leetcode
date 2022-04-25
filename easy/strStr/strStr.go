package main

import "fmt"

func main() {
	hay := "lelo"
	needle := "ll"
	fmt.Println(strStr(hay, needle))

}

func strStr(haystack string, needle string) int {
	n := len(needle)

	switch {
	case n == 0:
		return 0
	case n == 1:
		for i, i2 := range haystack {
			if uint8(i2) == needle[0] {
				return i
			}
		}
		return -1
	case n == len(haystack):
		if needle == haystack {
			return 0
		}
		return -1
	case n > len(haystack):
		return -1
	default:
		for i := 0; i < len(haystack)-n+1; i++ {
			if haystack[i] == needle[0] {
				if haystack[i:i+n] == needle {
					return i
				}
			}
		}
	}
	return -1
}
