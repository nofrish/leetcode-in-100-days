package _404

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

func sumOfLeftLeaves(root *TreeNode) (result int) {

	var dfs func(node *TreeNode, isLeft bool)
	dfs = func(node *TreeNode, isLeft bool) {
		if isLeft && node.Left == nil && node.Right == nil {
			result += node.Val
			return
		}
		if node.Left != nil {
			dfs(node.Left, true)
		}
		if node.Right != nil {
			dfs(node.Right, false)
		}
	}
	dfs(root, false)
	return
}
