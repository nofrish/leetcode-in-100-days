package _298

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestConsecutive_(root *TreeNode) (result int) {

	var dfs func(node *TreeNode, last int, count int)
	dfs = func(node *TreeNode, last int, count int) {
		if node == nil {
			return
		}
		if node.Val == last+1 {
			count++
			result = max(result, count)
		} else {
			count = 1
		}

		result = 1
		dfs(node.Left, node.Val, count)
		dfs(node.Right, node.Val, count)
	}

	dfs(root.Left, root.Val, 1)
	dfs(root.Right, root.Val, 1)
	return
}
