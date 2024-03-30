package main

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	matched := make([]int, 128)
	for _, val := range t {
		matched[val]++
	}

	left, right, valid, start, len := 0, 0, 0, 0, m+1
	for right < m {
		if matched[s[right]] > 0 {
			valid++
		}
		matched[s[right]]--
		right++
		for valid == n {
			if right-left < len {
				start = left
				len = right - left
			}
			if matched[s[left]] == 0 {
				valid--
			}
			matched[s[left]]++
			left++
		}
	}

	if len == m+1 {
		return ""
	}

	return s[start : start+len]
}
