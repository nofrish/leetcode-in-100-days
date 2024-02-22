package _094

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal_(root *TreeNode) (result []int) {

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			result = append(result, node.Val)
			dfs(node.Right)
		}
	}

	dfs(root)
	return
}
