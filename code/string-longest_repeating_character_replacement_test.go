package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		expect int
	}{
		// testcases from: https://leetcode.com/problems/longest-repeating-character-replacement/
		{s: "ABAB", k: 2, expect: 4},
		{s: "AABABBA", k: 1, expect: 4},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.s), func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("characterReplacement(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.expect)
			}
		})
	}
}
