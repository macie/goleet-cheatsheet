package main

// type ListNode struct {
//     Val int
//     Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil {
		if n > 0 {
			n--
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
