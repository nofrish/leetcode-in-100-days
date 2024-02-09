package _298

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

func longestConsecutive(root *TreeNode) int {

	var dfs func(node *TreeNode, parent int, count int)
	result := 0
	dfs = func(node *TreeNode, parent int, count int) {
		if node == nil {
			return
		}
		if node.Val == parent+1 {
			count++
		} else {
			count = 1
		}
		if count > result {
			result = count
		}

		dfs(node.Left, node.Val, count)
		dfs(node.Right, node.Val, count)
	}
	dfs(root, math.MaxInt, 0)
	return result
}
