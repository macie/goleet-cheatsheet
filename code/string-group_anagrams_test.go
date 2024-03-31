package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		// testcases from: https://leetcode.com/problems/group-anagrams/
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{strs: []string{""}, want: [][]string{{""}}},
		{strs: []string{"a"}, want: [][]string{{"a"}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.strs), func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) > len(got[j])
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(strs) = %v, want %v", got, tt.want)
			}
		})
	}
}
