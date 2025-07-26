package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	5
//
// 4        6
//
//	3     7
func isValidBST(root *TreeNode) bool {

	var fn func(r *TreeNode, mx, mn int) bool
	fn = func(r *TreeNode, mx, mn int) bool {
		if r == nil {
			return true
		}
		if r.Val >= mx || r.Val <= mn {
			return false
		}

		return fn(r.Left, r.Val, mn) && fn(r.Right, mx, r.Val)

	}
	return fn(root.Left, root.Val, math.MinInt32-1) && fn(root.Right, math.MaxInt32+1, root.Val)
}

func maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}
