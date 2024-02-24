package _530

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference_(root *TreeNode) (result int) {
	last := -1000000

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		result = min(result, int(math.Abs(float64(node.Val)-float64(last))))
		last = node.Val
		dfs(node.Right)
	}

	result = math.MaxInt
	dfs(root)
	return
}
