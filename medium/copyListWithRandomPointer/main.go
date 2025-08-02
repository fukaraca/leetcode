package main

func main() {
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var copier func(h *Node) *Node
	copier = func(h *Node) *Node {
		if h == nil {
			return nil
		}
		r := &Node{Val: h.Val}
		m[h] = r
		r.Next = copier(h.Next)
		if h.Random != nil {
			r.Random = m[h.Random]
		}
		return r
	}
	return copier(head)

}
