package _501

import "math"

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

func findMode(root *TreeNode) []int {

	result, maxCount := make([]int, 0), 0
	last, count := math.MinInt, 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if node.Val == last {
			count++
		} else {
			last, count = node.Val, 1
		}

		if count > maxCount {
			result, maxCount = []int{node.Val}, count
		} else if count == maxCount {
			result = append(result, node.Val)
		}

		dfs(node.Right)
	}

	dfs(root)
	return result
}
