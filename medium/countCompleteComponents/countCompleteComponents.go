package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5},
	}
	fmt.Println(countCompleteComponents(n, edges))
	edges = [][]int{
		{0, 1}, {0, 2}, {1, 2},
	}
	fmt.Println(countCompleteComponents(n, edges))
}

func countCompleteComponents(n int, edges [][]int) int {
	tree := make([]uint64, n)
	for i := range n {
		tree[i] = 1 << i // singular vertex is complete already
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		tree[b] = tree[b] | 1<<a // adjacent vertices have same bits
		tree[a] = tree[a] | 1<<b
	}
	counter := make(map[uint64]int)
	for _, u := range tree {
		counter[u]++ // count same signatures(bits)
	}
	var out int
	for u, i := range counter {
		if bits.OnesCount64(u) == i { // these vertices count should match otherwise it is not complete
			out++
		}
	}
	return out
}
