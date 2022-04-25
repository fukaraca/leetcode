package leafNodes

import "reflect"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(getList(root1), getList(root2))
}

func getList(root *TreeNode) []int {
	list := []int{}
	var travers func(*TreeNode)

	travers = func(node *TreeNode) {

		if node.Left != nil {
			travers(node.Left)
		}

		if node.Right != nil { //third case
			travers(node.Right)
		}
		if node.Left == nil && node.Right == nil { //base case
			list = append(list, node.Val)
		}
	}
	travers(root)
	return list
}
