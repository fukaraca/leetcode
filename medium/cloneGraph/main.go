package main

func main() {
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]*Node)
	var bfs func(n *Node) *Node
	bfs = func(n *Node) *Node {
		if v, ok := visited[n.Val]; ok {
			return v
		}
		out := &Node{Val: n.Val}
		visited[n.Val] = out
		for _, neighbor := range n.Neighbors {
			out.Neighbors = append(out.Neighbors, bfs(neighbor))
		}
		return out
	}
	return bfs(node)

}
