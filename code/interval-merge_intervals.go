package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := len(merged) - 1
		if merged[last][1] >= intervals[i][0] {
			merged[last][1] = max(
				merged[last][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
