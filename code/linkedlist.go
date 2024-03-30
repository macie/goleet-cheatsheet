package main

import (
	"fmt"
	"strings"
)

// ListNode represents a linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewCyclicLinkedList creates a linked list with an optional cycle between
// last node and node at `lastPtr` index. If `lastPtr` is -1, there is no cycle.
func NewCyclicLinkedList(vals []int, lastPtr int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	pointers := make([]*ListNode, len(vals))

	head := &ListNode{Val: vals[0]}
	curr := head
	pointers[0] = curr

	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
		pointers[i] = curr
	}

	if lastPtr >= 0 {
		curr.Next = pointers[lastPtr]
	}

	return head
}

func (l *ListNode) String() string {
	maxDepth := 6
	nodes := make([]string, 0, maxDepth)

	depth := 1
	curr := l
	for curr != nil && depth < maxDepth {
		nodes = append(nodes, fmt.Sprintf("%v", curr.Val))
		curr = curr.Next
		depth++
	}
	if depth == maxDepth {
		nodes = append(nodes, "...")
	}
	return "[" + strings.Join(nodes, "->") + "]"
}
