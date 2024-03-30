package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		heights [][]int
		want    [][]int
	}{
		// testcases from: https://leetcode.com/problems/pacific-atlantic-water-flow/
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{heights: [][]int{{1}}, want: [][]int{{0, 0}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.heights), func(t *testing.T) {
			got := pacificAtlantic(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pacificAtlantic(heights) = %v, want %v", got, tt.want)
			}
		})
	}
}
