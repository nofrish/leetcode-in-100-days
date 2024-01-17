package lintcode_391

import (
	"math"
	"sort"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

func CountOfAirplanes_(airplanes []*Interval) int {

	type Node struct {
		Time     int
		UpOrDown int
	}

	nodes := make([]Node, 0, len(airplanes)*2)
	for _, airplane := range airplanes {
		nodes = append(nodes, Node{airplane.Start, 1})
		nodes = append(nodes, Node{airplane.End, -1})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Time == nodes[j].Time {
			return nodes[i].UpOrDown < nodes[j].UpOrDown
		}
		return nodes[i].Time < nodes[j].Time
	})

	count, max := 0, 0
	for _, node := range nodes {
		count += node.UpOrDown
		max = int(math.Max(float64(max), float64(count)))
	}
	return max
}
