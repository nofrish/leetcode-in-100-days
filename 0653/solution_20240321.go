package _653

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

func findTarget(root *TreeNode, k int) bool {

	exists := make(map[int]struct{})
	found := false

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		other := k - node.Val
		if _, ok := exists[other]; ok {
			found = true
		}
		exists[node.Val] = struct{}{}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return found
}
