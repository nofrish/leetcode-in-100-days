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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	var dfs func(*TreeNode) bool

	prev := math.MinInt
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if dfs(node.Left) == false {
			return false
		}
		if prev >= node.Val {
			return false
		}
		prev = node.Val
		return dfs(node.Right)
	}

	return dfs(root)
}
