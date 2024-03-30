package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		// testcases from: https://leetcode.com/problems/reverse-linked-list/
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}},
		},
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
		},
		{
			head: &ListNode{},
			want: &ListNode{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.head), func(t *testing.T) {
			got := reverseList(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList(head) = %v, want %v", got, tt.want)
			}
		})
	}
}
