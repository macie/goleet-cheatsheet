package main

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		// testcases from: https://leetcode.com/problems/linked-list-cycle/
		{head: NewCyclicLinkedList([]int{3, 2, 0, -4}, 1), want: true},
		{head: NewCyclicLinkedList([]int{1, 2}, -1), want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.head), func(t *testing.T) {
			got := hasCycle(tt.head)
			if got != tt.want {
				t.Errorf("hasCycle(head) = %v, want %v", got, tt.want)
			}
		})
	}
}
