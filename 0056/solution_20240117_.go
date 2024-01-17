package _056

import "sort"

func merge__(intervals [][]int) [][]int {

	type Node struct {
		X          int
		StartOrEnd int
	}

	nodes := make([]Node, 0, len(intervals)*2)
	for _, interval := range intervals {
		nodes = append(nodes, Node{interval[0], 1})
		nodes = append(nodes, Node{interval[1], -1})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].X == nodes[j].X {
			return nodes[i].StartOrEnd > nodes[j].StartOrEnd
		}
		return nodes[i].X < nodes[j].X
	})

	result := make([][]int, 0)
	start, count := -1, 0
	for _, node := range nodes {
		count += node.StartOrEnd
		if count > 0 && start < 0 {
			start = node.X
		}
		if count == 0 {
			result = append(result, []int{start, node.X})
			start = -1
		}
	}

	return result
}
