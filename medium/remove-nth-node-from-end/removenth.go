package remove_nth_node_from_end

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}

	nTh := findNth(temp, n)
	nTh.Next = nTh.Next.Next
	return temp.Next

}

func findNth(temp *ListNode, n int) *ListNode {
	pioneer := temp
	rear := temp
	for i := 0; i < n; i++ {
		pioneer = pioneer.Next
	}
	for pioneer.Next != nil {
		rear = rear.Next
		pioneer = pioneer.Next
	}
	return rear

}
