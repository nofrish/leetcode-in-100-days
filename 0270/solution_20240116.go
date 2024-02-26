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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) (result int) {

	min := math.MaxFloat64
	for root != nil {
		abs := math.Abs(float64(root.Val) - target)
		if (min > abs) || (min == abs && root.Val < result) {
			min = abs
			result = root.Val
		}
		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return
}
