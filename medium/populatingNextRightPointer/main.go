package main

func main() {
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	leveler := make(map[int]*Node)
	var dfs func(r *Node, l int)
	dfs = func(r *Node, l int) {
		if r == nil {
			return
		}
		dfs(r.Right, l+1)
		dfs(r.Left, l+1)
		r.Next = leveler[l]
		leveler[l] = r
	}
	dfs(root, 0)
	return root

}
