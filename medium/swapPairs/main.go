package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
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

	//fmt.Println(swapPairs(list.head))
	printPairs(swapPairs(list.head))

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	head = curr.Next
	var prev *ListNode = nil
	for curr != nil && curr.Next != nil {
		next := curr.Next.Next //get next pairs first
		curr.Next.Next = curr  //bind current pairs second to current pairs first
		if prev != nil {       //bind previous pair to current pairs second
			prev.Next = curr.Next
		}
		curr.Next = next //bind current pairs first to next pair
		prev = curr      //back up current updated pair
		curr = curr.Next //set current pair to next pair
	}
	return head
}
