package main

import "reflect"

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	sMap := make(map[string]int)
	tMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		sMap[string(s[i])]++
		tMap[string(t[i])]++
	}

	return reflect.DeepEqual(sMap, tMap)
}
