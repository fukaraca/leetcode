package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	head := ret
	headed := true
	residue := 0
	for {
		sum := 0
		if l1 == nil {
			sum = l2.Val + residue
		} else if l2 == nil {
			sum = l1.Val + residue
		} else {
			sum = l1.Val + l2.Val + residue
		}

		if sum > 9 {
			ret.Val = sum % 10
			residue = 1
		} else {
			ret.Val = sum
			residue = 0
		}

		if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
			if residue > 0 {
				ret.Next = new(ListNode)
				ret.Next.Val = residue
			}
			break
		}

		if headed {
			headed = false
			ret.Next = new(ListNode)
			head = ret
		} else {
			ret.Next = new(ListNode)
		}
		if l1 == nil {
			l2 = l2.Next
		} else if l2 == nil {
			l1 = l1.Next
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}

		ret = ret.Next
	}
	return head
}
