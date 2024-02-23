package _448

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes_(root *TreeNode) (result int) {

	var dfs func(node *TreeNode, cMax int) // cMax is current max
	dfs = func(node *TreeNode, cMax int) {
		if node == nil {
			return
		}
		if node.Val >= cMax {
			result += 1
			cMax = node.Val
		}
		dfs(node.Left, cMax)
		dfs(node.Right, cMax)
	}

	dfs(root, math.MinInt)
	return
}
