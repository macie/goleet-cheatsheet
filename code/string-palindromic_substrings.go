package main

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += countPalindrome(s, i, i)
		count += countPalindrome(s, i, i+1)
	}
	return count
}

func countPalindrome(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}
