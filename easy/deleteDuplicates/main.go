package deleteDuplicates

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}

	for head.Next != nil {

		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		} else {
			head = head.Next
		}

	}
	return dummy.Next
}
