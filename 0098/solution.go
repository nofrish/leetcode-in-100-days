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
