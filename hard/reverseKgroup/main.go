package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := &LinkedList{}

	for _, num := range a {
		temp := &ListNode{
			Val:  num,
			Next: nil,
		}
		if list.head == nil {
			list.head = temp
			continue
		}
		l := list.head
		for l.Next != nil {
			l = l.Next
		}

		l.Next = temp
	}
	temp := list.head
	for temp != nil {
		temp = temp.Next
	}

	printPairs(reverseKGroup(list.head, 11))

}

func printPairs(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	cur := head
	for i := 0; i < k-1; i++ {
		head = head.Next
	}
	var prev *ListNode = nil

	for {
		stack := []*ListNode{}
		first := cur                             //get first
		for i := 0; i < k-1 && cur != nil; i++ { //insert middles to stack
			if i >= 1 && i <= k-2 {
				stack = append(stack, cur)
			}
			cur = cur.Next //at tail
		}
		if cur == nil { //ensure valid k group
			break
		}
		first.Next = cur.Next //bind first to next group
		if prev != nil {      //bind prev to tail
			prev.Next = cur
		}
		for i := len(stack) - 1; i >= 0; i-- { //bind tail to middles
			cur.Next = stack[i]
			cur = cur.Next
		}
		cur.Next = first //bind middles to first
		prev = cur.Next  //save previous
		cur = prev.Next  //set current
	}

	return head
}
