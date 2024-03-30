package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// testcases from: https://leetcode.com/problems/container-with-most-water/
		{nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{nums: []int{1, 1}, want: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := maxArea(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxArea(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
