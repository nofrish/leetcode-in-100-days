package _110

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

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) (bool, int)
	dfs = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}

		lBalanced, h1 := dfs(root.Left)
		if !lBalanced {
			return false, 0
		}

		rBalanced, h2 := dfs(root.Right)
		if !rBalanced {
			return false, 0
		}

		if math.Abs(float64(h1)-float64(h2)) > 1 {
			return false, 0
		}
		return true, max(h1, h2) + 1
	}

	balanced, _ := dfs(root)
	return balanced
}
