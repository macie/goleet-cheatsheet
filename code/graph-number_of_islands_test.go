package main

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		// testcases from: https://leetcode.com/problems/number-of-islands/
		{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, want: 1},
		{grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.grid), func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.want {
				t.Errorf("numIslands(grid) = %v, want %v", got, tt.want)
			}
		})
	}
}
