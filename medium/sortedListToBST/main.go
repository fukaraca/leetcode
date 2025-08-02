package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var list []int

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	var makeBalanced func(ls []int) *TreeNode
	makeBalanced = func(ls []int) *TreeNode {
		if len(ls) == 0 {
			return nil
		}
		m := len(ls) / 2
		node := &TreeNode{Val: ls[m]}
		node.Left = makeBalanced(ls[:m])
		node.Right = makeBalanced(ls[m+1:])
		return node
	}
	return makeBalanced(list)
}
