package main

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		// testcases from: https://leetcode.com/problems/minimum-window-substring/
		{s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
		{s: "a", t: "a", want: "a"},
		{s: "a", t: "aa", want: ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.s, tt.t), func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(s, t) = %v, want %v", got, tt.want)
			}
		})
	}
}
