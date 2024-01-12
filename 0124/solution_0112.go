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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	_, max := traverse(root)
	return max
}

func traverse(node *TreeNode) (pathMax, max int) {
	if node == nil {
		return 0, math.MinInt
	}

	lcPathMax, lcMax := traverse(node.Left)
	rcPathMax, rcMax := traverse(node.Right)

	lcPathMax = int(math.Max(0.0, float64(lcPathMax)))
	rcPathMax = int(math.Max(0.0, float64(rcPathMax)))

	if lcPathMax > rcPathMax {
		pathMax = lcPathMax + node.Val
	} else {
		pathMax = rcPathMax + node.Val
	}

	max = int(math.Max(math.Max(float64(lcMax), float64(rcMax)), float64(lcPathMax+node.Val+rcPathMax)))
	return
}
