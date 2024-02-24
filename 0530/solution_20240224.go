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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	result := math.MaxInt

	var dfs func(node *TreeNode, lower, upper int)
	dfs = func(node *TreeNode, lower, upper int) {
		if node == nil {
			return
		}
		result = min(result, abs(node.Val-lower), abs(node.Val-upper))
		dfs(node.Left, lower, node.Val)
		dfs(node.Right, node.Val, upper)
	}

	dfs(root, -1000000, 1000000)
	return result
}
