package _270

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func closestValue_(root *TreeNode, target float64) (result int) {

	min := math.MaxFloat64

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		abs := math.Abs(float64(node.Val) - target)
		if abs < min || (abs == min && node.Val < result) {
			min = abs
			result = node.Val
		}
		if float64(node.Val) > target {
			dfs(node.Left)
		} else {
			dfs(node.Right)
		}
	}

	dfs(root)
	return
}
