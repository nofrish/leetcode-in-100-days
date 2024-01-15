package _314

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Depth int
	Val   int
}

func verticalOrder(root *TreeNode) (result [][]int) {

	columns := make(map[int][]Node, 0)
	minColumnNum := 0

	var dfs func(node *TreeNode, depth, columnNum int)
	dfs = func(node *TreeNode, depth, columnNum int) {
		if node != nil {
			if minColumnNum > columnNum {
				minColumnNum = columnNum
			}
			if column, ok := columns[columnNum]; ok {
				columns[columnNum] = append(column, Node{depth, node.Val})
			} else {
				columns[columnNum] = []Node{Node{depth, node.Val}}
			}
			dfs(node.Left, depth+1, columnNum-1)
			dfs(node.Right, depth+1, columnNum+1)
		}
	}
	dfs(root, 0, 0)

	for column := columns[minColumnNum]; len(column) > 0; {

		sort.SliceStable(column, func(i, j int) bool {
			return column[i].Depth < column[j].Depth
		})
		vals := make([]int, len(column))
		for i, node := range column {
			vals[i] = node.Val
		}
		result = append(result, vals)

		minColumnNum++
		column = columns[minColumnNum]
	}
	return
}
