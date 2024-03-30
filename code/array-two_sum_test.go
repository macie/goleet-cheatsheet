package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		// testcases from: https://leetcode.com/problems/two-sum
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%v, %d)", tt.nums, tt.target), func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(nums, target) = %v, want %v", got, tt.want)
			}
		})
	}
}
