package _229

import (
	"math"
	"sort"
)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {

	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	for i, j := 0, 0; i < len(slots1) && j < len(slots2); {

		s1, s2 := slots1[i], slots2[j]

		// 有重合情况，直接计算重合长度，并更新Index
		lower := int(math.Max(float64(s1[0]), float64(s2[0])))
		upper := int(math.Min(float64(s1[1]), float64(s2[1])))
		if upper-lower >= duration {
			return []int{lower, lower + duration}
		}

		if s1[1] > s2[1] {
			j++
		} else {
			i++
		}
	}

	return []int{}
}
