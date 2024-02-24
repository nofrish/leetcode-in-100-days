package _098

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST__(root *TreeNode) bool {

	var dfs func(node *TreeNode, lower, upper int) bool
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
