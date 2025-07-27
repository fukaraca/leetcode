package main

import "math"

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(r *TreeNode) (int, bool)
	dfs = func(r *TreeNode) (int, bool) {
		if r == nil {
			return 0, true
		}

		ll, ok := dfs(r.Left)
		if !ok {
			return 0, false
		}
		lr, ok := dfs(r.Right)
		if !ok || math.Abs(float64(ll-lr)) > 1 {
			return 0, false
		}
		if ll > lr {
			return ll + 1, true
		}
		return lr + 1, true
	}
	l, ok := dfs(root.Left)
	if !ok {
		return false
	}
	r, ok := dfs(root.Right)
	if !ok {
		return false
	}
	return math.Abs(float64(l-r)) < 2
}
