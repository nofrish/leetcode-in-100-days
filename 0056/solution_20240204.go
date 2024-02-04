package _056

import "sort"

func merge___(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := result[len(result)-1]
		if last[1] >= interval[0] {
			last[1] = max(interval[1], last[1])
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func merge____(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	slow, fast := 0, 0
	for fast++; fast < len(intervals); fast++ {
		if intervals[slow][1] >= intervals[fast][0] {
			intervals[slow][1] = max(intervals[slow][1], intervals[fast][1])
		} else {
			slow++
			intervals[slow] = intervals[fast]
		}
	}
	return intervals[:slow+1]
}
