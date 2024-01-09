package _252

import "sort"

type Node struct {
	Time  int
	Count int
}

func canAttendMeetings(intervals [][]int) bool {

	nodes := make([]Node, 0, len(intervals)*2)
	for _, interval := range intervals {
		nodes = append(nodes, Node{interval[0], 1})
		nodes = append(nodes, Node{interval[1], -1})
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Time == nodes[j].Time {
			return nodes[i].Count < nodes[j].Count
		}
		return nodes[i].Time < nodes[j].Time
	})

	count := 0
	for _, node := range nodes {
		count += node.Count
		if count > 1 {
			return false
		}
	}
	return true
}
