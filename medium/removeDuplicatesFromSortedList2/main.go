package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	m := make(map[int]int)
	var next func(h *ListNode) *ListNode
	next = func(h *ListNode) *ListNode {
		if h == nil {
			return nil
		}
		m[h.Val]++
		n := next(h.Next)
		if m[h.Val] > 1 {
			return n
		}
		h.Next = n
		return h
	}
	return next(head)
}
