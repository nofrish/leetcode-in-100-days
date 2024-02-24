package _700

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST_(root *TreeNode, val int) (result *TreeNode) {

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if dfs(node.Left) || dfs(node.Right) {
			return true
		}
		if node.Val == val {
			result = node
			return true
		}
		return false
	}
	if !dfs(root) {
		result = nil
	}
	return
}

func searchBST__(root *TreeNode, val int) (result *TreeNode) {

	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
