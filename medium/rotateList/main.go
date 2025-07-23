package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	var stack []*ListNode
	n := head

	for n != nil {
		t := n
		stack = append(stack, t) //1-2-3-4-5 >>> 5-1-2-3-4
		n = n.Next
	}
	l := len(stack)
	k %= l
	for i := 0; i < k; i++ {
		stack = append(stack[l-1:], stack[:l-1]...)
		stack[0].Next = stack[1]
		stack[l-1].Next = nil
	}

	return stack[0]
}
