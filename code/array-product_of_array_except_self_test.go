package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		// testcases from: https://leetcode.com/problems/product-of-array-except-self/
		{nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(nums) = %v, want %v", got, tt.want)
			}
		})
	}
}
