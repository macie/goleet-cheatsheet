package main

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	count := make(map[byte]int, len(s))
	for i := range s {
		count[s[i]]++
		count[t[i]]--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
