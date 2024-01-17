package _253

import "sort"

func minMeetingRooms(intervals [][]int) int {
	type Node struct {
		Time       int
		StartOrEnd int
	}

	nodes := make([]Node, 0, len(intervals)*2)
	for _, interval := range intervals {
		nodes = append(nodes, Node{interval[0], 1})
		nodes = append(nodes, Node{interval[1], -1})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Time == nodes[j].Time {
			return nodes[i].StartOrEnd < nodes[j].StartOrEnd
		}
		return nodes[i].Time < nodes[j].Time
	})

	max, count := 0, 0
	for _, node := range nodes {
		count += node.StartOrEnd
		if max < count {
			max = count
		}
	}
	return max
}
