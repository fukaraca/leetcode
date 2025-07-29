package main

import "fmt"

func main() {
	fmt.Println(reverse([][]int{{1}, {2, 3}, {4, 5}}))
	fmt.Println(reverse([][]int{{1}}))
	fmt.Println(reverse([][]int{{1}, {2, 3}, {4, 5}, {6, 7}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var out [][]int
	type qNode struct {
		node  *TreeNode
		level int
	}
	var q []*qNode

	q = append(q, &qNode{root, 0})

	for len(q) > 0 {
		r := q[0]
		q = q[1:]
		if r.node == nil {
			continue
		}

		if len(out) == r.level {
			out = append(out, []int{})
		}
		out[r.level] = append(out[r.level], r.node.Val)
		q = append(q, &qNode{r.node.Left, r.level + 1}, &qNode{r.node.Right, r.level + 1})
	}

	return reverse(out)
}

func reverse(out [][]int) [][]int {
	l := len(out)
	if l < 2 {
		return out
	}
	for i, j := 0, l-1; i < l/2 && j >= 0; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
