package main

import "sort"

func threeSum(nums []int) [][]int {
	triplets := make([][]int, 0, len(nums))

	sort.Ints(nums)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c < 0 {
				l++
			} else if a+b+c == 0 {
				triplets = append(triplets, []int{a, b, c})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
				r--
			} else {
				r--
			}
		}
	}

	return triplets
}
