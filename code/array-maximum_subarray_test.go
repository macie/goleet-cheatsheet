package main

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// testcases from: https://leetcode.com/problems/maximum-subarray/
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{nums: []int{1}, want: 1},
		{nums: []int{5, 4, -1, 7, 8}, want: 23},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
