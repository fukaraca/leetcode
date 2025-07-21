package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(sl, 4))
}

func topKFrequent(words []string, k int) []string {
	l := len(words)
	m := make(map[string]int, l)
	out := make([]string, 0)

	for _, w := range words {
		if _, ok := m[w]; ok {
			m[w]--
		} else {
			m[w] = l
		}
	}
	m2 := make(map[int][]string)
	for s, i := range m {
		m2[i] = append(m2[i], s)
	}
	for i := 0; i <= l && k > 0; i++ {
		if ln := len(m2[i]); ln > 0 {
			sl := m2[i]
			sort.Slice(sl, func(i, j int) bool {
				return sl[i] < sl[j]
			})
			if k >= ln {
				k -= ln
			} else {
				out = append(out, sl[:k]...)
				break
			}

			out = append(out, sl...)
		}
	}
	return out
}
