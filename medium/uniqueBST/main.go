package main

import "fmt"

func main() {
}

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	type a struct {
		s, e int
	}

	m := make(map[a]*int)

	var tree func(a) *int

	tree = func(ind a) *int {
		if ind.s > ind.e {
			r := 1
			return &r
		}
		if v, ok := m[ind]; ok {
			return v
		}

		var out int
		for i := ind.s; i <= ind.e; i++ {
			leftNodes := tree(a{ind.s, i - 1})
			rightNodes := tree(a{i + 1, ind.e})
			fmt.Println(*leftNodes, *rightNodes)
			out += *leftNodes * *rightNodes
		}

		m[ind] = &out
		return &out
	}
	out := tree(a{1, n})
	return *out
}
