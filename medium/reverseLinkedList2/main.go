package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var stack []*ListNode

	var fn func(h *ListNode, l, r int) *ListNode
	fn = func(h *ListNode, l, r int) *ListNode {
		if r < 0 {
			return h
		}
		if l <= 0 && r >= 0 {
			stack = append(stack, h)
		}
		ret := fn(h.Next, l-1, r-1)

		if r < len(stack) {
			stack[r].Next = ret
			return stack[r]
		}
		h.Next = ret
		return h
	}
	return fn(head, left-1, right-1)
}
