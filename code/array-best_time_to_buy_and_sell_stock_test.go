package main

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		// testcases from: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.prices), func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(prices) = %v, want %v", got, tt.want)
			}
		})
	}
}
