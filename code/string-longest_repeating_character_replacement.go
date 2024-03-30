package main

func characterReplacement(s string, k int) int {
	m := map[byte]int{}
	maxLen := 0
	maxRepeat := byte(0)
	left, right := 0, 0
	for right < len(s) {
		char := s[right]
		m[char]++

		count := m[char]
		if maxRepeat == 0 || m[maxRepeat] < count {
			maxRepeat = char
		}

		if right-left+1-m[maxRepeat] > k {
			m[s[left]]--
			left++
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
		right++
	}

	return maxLen
}
