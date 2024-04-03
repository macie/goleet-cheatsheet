package main

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		// testcases from: https://leetcode.com/problems/valid-anagram/
		{s: "anagram", t: "nagaram", want: true},
		{s: "rat", t: "car", want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.s, tt.t), func(t *testing.T) {
			got := validAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("ValidAnagram(s, t) = %v, want %v", got, tt.want)
			}
		})
	}
}
