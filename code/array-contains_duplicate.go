package main

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool)

	for _, n := range nums {
		if exist[n] {
			return true
		}
		exist[n] = true
	}

	return false
}
