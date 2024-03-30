package main

func lengthOfLongestSubstring(s string) int {
	var longest, left, right int
	seen := make([]bool, 128) // 128 ASCII characters

	for right < len(s) {
		if !seen[s[right]-'a'] {
			seen[s[right]-'a'] = true
			right++
			longest = max(longest, right-left)
		} else {
			seen[s[left]-'a'] = false
			left++
		}
	}
	return longest
}
