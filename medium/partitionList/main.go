package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var stack []*ListNode
	var traverse func(h *ListNode) *ListNode
	traverse = func(h *ListNode) *ListNode {
		if h == nil {
			return nil
		}
		n := traverse(h.Next)
		if h.Val < x {
			stack = append(stack, h)
			return n
		}
		h.Next = n
		return h
	}
	head = traverse(head)
	for i := 0; i < len(stack); i++ {
		stack[i].Next = head
		head = stack[i]
	}
	return head
}
