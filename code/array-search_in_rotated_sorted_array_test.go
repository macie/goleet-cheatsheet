package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		// testcases from: https://leetcode.com/problems/search-in-rotated-sorted-array/
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{1}, target: 0, want: -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search(nums, target) = %v, want %v", got, tt.want)
			}
		})
	}
}
