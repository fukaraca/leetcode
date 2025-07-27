package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	var trav func(l, r *TreeNode) bool
	trav = func(l, r *TreeNode) bool {
		if l == r {
			return true
		}
		if l != nil && r == nil || l == nil && r != nil || l.Val != r.Val {
			return false
		}
		return trav(l.Left, r.Right) && trav(l.Right, r.Left)
	}
	return trav(root.Left, root.Right)
}
