package _759

import "sort"

/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedules [][]*Interval) []*Interval {

	type Node struct {
		Time       int
		BusyOrFree int
	}

	nodes := make([]Node, 0)
	for _, schedule := range schedules {
		for _, interval := range schedule {
			nodes = append(nodes, Node{interval.Start, 1})
			nodes = append(nodes, Node{interval.End, -1})
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Time == nodes[i].Time {
			return nodes[i].BusyOrFree > nodes[j].BusyOrFree
		}
		return nodes[i].Time < nodes[j].Time
	})

	result := make([]*Interval, 0)
	start, count := -1, 0
	for _, node := range nodes {
		count += node.BusyOrFree
		if count == 0 {
			start = node.Time
		}
		if count > 0 && start > 0 {
			result = append(result, &Interval{start, node.Time})
			start = -1
		}
	}

	return result
}
