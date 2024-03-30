package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		// testcases from: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			n:    2,
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}},
		},
		{
			head: &ListNode{Val: 1, Next: nil},
			n:    1,
			want: nil,
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			n:    1,
			want: &ListNode{Val: 1, Next: nil},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.head), func(t *testing.T) {
			got := removeNthFromEnd(tt.head, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd(head, n) = %v, want %v", got, tt.want)
			}
		})
	}
}
