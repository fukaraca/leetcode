package main

import (
	"container/heap"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type pq []*ListNode

func (r pq) Len() int {
	return len(r)
}
func (r pq) Less(i, j int) bool {
	return (r)[i].Val < (r)[j].Val
}

func (r pq) Swap(i, j int) {
	(r)[i], (r)[j] = (r)[j], (r)[i]
}

func (r *pq) Push(x interface{}) {
	*r = append(*r, x.(*ListNode))
}
func (r *pq) Pop() interface{} {

	n := len(*r)
	x := (*r)[n-1]

	*r = (*r)[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	pl := pq{}
	heap.Init(&pl)
	head := new(ListNode)
	dummy := head
	if len(lists) == 0 {
		return nil
	}

	for _, list := range lists {
		if list != nil {
			//push only heads. heads are already sorted ascending order beforehand
			heap.Push(&pl, list)
		}
	}
	for pl.Len() > 0 {
		//it returns smallest valued list always
		node := heap.Pop(&pl).(*ListNode)
		if node.Next != nil {
			//push next node to heap
			heap.Push(&pl, node.Next)
		}
		head.Next = node
		head = head.Next
	}
	return dummy.Next
}
