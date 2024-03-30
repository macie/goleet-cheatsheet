package main

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		// testcases from: https://leetcode.com/problems/valid-parentheses/
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(]", want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := validParentheses(tt.s)
			if got != tt.want {
				t.Errorf("ValidParentheses(s) = %v, want %v", got, tt.want)
			}
		})
	}
}
