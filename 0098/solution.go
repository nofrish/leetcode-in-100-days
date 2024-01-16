package _098

import "math"

// Solution One.

func _isValidBST(root *TreeNode) bool {

	var dfs func(node *TreeNode) bool

	prev := math.MinInt
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
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

// Solution Two.
func __isValidBST(root *TreeNode) bool {

	var dfs func(node *TreeNode, lowerBound, upperBound int) bool
	dfs = func(node *TreeNode, lowerBound, upperBound int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lowerBound {
			return false
		}
		if node.Val >= upperBound {
			return false
		}
		return dfs(node.Left, lowerBound, node.Val) && dfs(node.Right, node.Val, upperBound)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
