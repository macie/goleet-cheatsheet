package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		// testcases from: https://leetcode.com/problems/3sum/
		{nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{0, 1, 1}, want: [][]int{}},
		{nums: []int{0, 0, 0}, want: [][]int{{0, 0, 0}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := threeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
