package main

import "fmt"

func main() {
	n1 := []int{2, 4, 3}
	n2 := []int{5, 6, 4}
	l1, l2 := generate(n1), generate(n2)
	fmt.Println(addTwoNumbersTwo(l1, l2))

}

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

func addTwoNumbersTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{}
	var elde int
	t := out
	for l1 != nil || l2 != nil || elde != 0 {
		if l1 != nil {
			t.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t.Val += l2.Val
			l2 = l2.Next
		}
		t.Val += elde
		elde = t.Val / 10
		t.Val %= 10

		t.Next = func() *ListNode {
			if l1 != nil || l2 != nil || elde != 0 {
				return new(ListNode)
			}
			return nil
		}()
		t = t.Next
	}
	return out
}

func generate(list []int) *ListNode {
	var out *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		out = &ListNode{
			Val: list[i],
			Next: func() *ListNode {
				if out == nil {
					return nil
				}
				t := *out
				return &t
			}(),
		}
	}
	return out
}
