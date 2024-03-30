package main

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		// testcases from: https://leetcode.com/problems/contains-duplicate/
		{nums: []int{1, 2, 3, 1}, want: true},
		{nums: []int{1, 2, 3, 4}, want: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
