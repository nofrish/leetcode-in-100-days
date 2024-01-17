package _435

import (
	"sort"
)

func _eraseOverlapIntervals(intervals [][]int) (removed int) {

	// 先排序，让所有可能冲突的区间在一起
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 以下的逻辑为：
	// - 让cur指向第一个区间，并循环对比下一个区间和cur的关系
	// - 如果下一个区间和cur不重叠，则cur指向下一个区间
	// - 如果下一个区间和cur重叠，则removed++，同时让cur指向区间右边界更小的区间
	cur := intervals[0]
	for _, interval := range intervals[1:] {
		if cur[1] <= interval[0] { // 不重叠
			cur = interval
			continue
		}
		removed++
		if cur[1] > interval[1] { // 让cur指向右边界更小的区间，为后面的区间预留更多的位置
			cur = interval
		}
	}
	return
}
