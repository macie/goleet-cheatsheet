package main

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// testcases from: https://leetcode.com/problems/maximum-product-subarray/
		{nums: []int{2, 3, -2, 4}, want: 6},
		{nums: []int{-2, 0, -1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := maxProduct(tt.nums)
			if got != tt.want {
				t.Errorf("maxProduct(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
