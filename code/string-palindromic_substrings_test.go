package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// testcases from: https://leetcode.com/problems/palindromic-substrings/
		{s: "abc", want: 3},
		{s: "aaa", want: 6},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := countSubstrings(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubstrings(s) = %v, want %v", got, tt.want)
			}
		})
	}
}
