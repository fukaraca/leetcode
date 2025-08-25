package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	type a struct {
		s, e int
	}

	m := make(map[a][]*TreeNode)

	var gen func(ind a) []*TreeNode

	gen = func(ind a) []*TreeNode {
		if ind.s > ind.e {
			return []*TreeNode{nil}
		}
		if v, ok := m[ind]; ok {
			return v
		}
		var out []*TreeNode
		for i := ind.s; i <= ind.e; i++ {
			leftNode := gen(a{ind.s, i - 1})
			rightNode := gen(a{i + 1, ind.e})

			for _, nodel := range leftNode {
				for _, noder := range rightNode {
					out = append(out, &TreeNode{Val: i, Left: nodel, Right: noder})
				}
			}
		}
		m[ind] = out
		return out
	}

	return gen(a{1, n})
}
