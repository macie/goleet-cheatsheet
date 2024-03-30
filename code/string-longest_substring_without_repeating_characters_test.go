package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// testcases from: https://leetcode.com/problems/longest-substring-without-repeating-characters/
		{s: "abcabcbb", want: 3},
		{s: "bbbbb", want: 1},
		{s: "pwwkew", want: 3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
