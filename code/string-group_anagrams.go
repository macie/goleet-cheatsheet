package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sorted := sortLetters(str)
		m[sorted] = append(m[sorted], str)
	}

	result := make([][]string, 0)
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

func sortLetters(str string) string {
	s := []rune(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}
