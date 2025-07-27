package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return root != nil &&
		((targetSum-root.Val == 0 && root.Left == nil && root.Right == nil) ||
			hasPathSum(root.Left, targetSum-root.Val) ||
			hasPathSum(root.Right, targetSum-root.Val))
}
