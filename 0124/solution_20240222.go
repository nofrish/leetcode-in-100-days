package _124

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum_(root *TreeNode) int {
	result := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lc := max(0, dfs(node.Left))  // 如果左分支比0还小，直接砍断
		rc := max(0, dfs(node.Right)) // 如果右分支比0还小，直接砍断
		result = max(result, lc+node.Val+rc)
		return max(lc+node.Val, rc+node.Val)
	}
	dfs(root)
	return result
}
