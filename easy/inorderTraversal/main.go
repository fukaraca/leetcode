package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var out []int
	var fn func(*TreeNode)
	fn = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			fn(r.Left)
		}
		out = append(out, r.Val)
		if r.Right != nil {
			fn(r.Right)
		}
	}

	fn(root)
	return out
}
