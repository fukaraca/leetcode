package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring2("abcacbcd"))
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	max := 1
	substr := make(map[int32]int)
	if s == "" {
		return 0
	}
	for i, l := range s {
		if ind, ok := substr[l]; ok {
			if i-start > max {
				max = i - start
			}
			if ind >= start {
				start = ind + 1
			}
			substr[l] = i
		} else {
			substr[l] = i
		}

	}
	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	var maks, base int
	m := make(map[uint8]int)
	for i := range s {
		if v, ok := m[s[i]]; ok {
			for ; base <= v; base++ {
				delete(m, s[base])
			}
		}

		m[s[i]] = i
		if maks < len(m) {
			maks = len(m)
		}
	}
	return maks
}
