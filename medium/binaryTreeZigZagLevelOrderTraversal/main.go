package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		if r.level%2 == 0 {
			out[r.level] = append(out[r.level], r.node.Val)
		} else {
			out[r.level] = append(out[r.level], 0)
			copy(out[r.level][1:], out[r.level][:len(out[r.level])-1])
			out[r.level][0] = r.node.Val
		}

		q = append(q, &qNode{r.node.Left, r.level + 1}, &qNode{r.node.Right, r.level + 1})
	}
	return out
}
