package _288

import "sort"

func removeCoveredIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	remains := 1

	for _, interval := range intervals {
		if interval[1] > cur[1] {
			remains++
			cur = interval
		}
	}

	return remains
}
