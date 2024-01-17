package _252

import "sort"

func canAttendMeetings_(intervals [][]int) bool {

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

	count := 0
	for _, node := range nodes {
		count += node.StartOrEnd
		if count > 1 {
			return false
		}
	}
	return true
}
