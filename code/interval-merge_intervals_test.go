package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		// testcases from: https://leetcode.com/problems/merge-intervals/
		{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			got := merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge(intervals) = %v, want %v", got, tt.want)
			}
		})
	}
}
