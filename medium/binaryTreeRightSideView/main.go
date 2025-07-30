package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	lm := make(map[int]int)
	var dfs func(r *TreeNode, l int)
	dfs = func(r *TreeNode, l int) {
		if r == nil {
			return
		}
		lm[l] = r.Val
		dfs(r.Left, l+1)
		dfs(r.Right, l+1)
	}
	dfs(root, 0)
	out := make([]int, len(lm))
	for i, i2 := range lm {
		out[i] = i2
	}
	return out
}
