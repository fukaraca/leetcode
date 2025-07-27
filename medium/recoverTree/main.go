package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 			3
// 		2		4
// 	5				1

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	var fn func(r *TreeNode)
	fn = func(r *TreeNode) {
		if r == nil {
			return
		}
		fn(r.Left)
		if prev != nil && r.Val <= prev.Val {
			if first == nil {
				first = prev
			}
			second = r
		}
		prev = r
		fn(r.Right)
	}
	fn(root)
	first.Val, second.Val = second.Val, first.Val
}
