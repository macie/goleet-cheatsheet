package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// testcases from: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
		{nums: []int{3, 4, 5, 1, 2}, want: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{nums: []int{11, 13, 15, 17}, want: 11},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findMin(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMin(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
