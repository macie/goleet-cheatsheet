package main

import "sort"

func merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0, size)
	merged = append(merged, intervals[0])
	for _, interval := range intervals[1:] {
		latest := merged[len(merged)-1]
		if interval[0] > latest[1] {
			merged = append(merged, interval)
		} else {
			latest[1] = max(interval[1], latest[1])
		}
	}

	return merged
}
