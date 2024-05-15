package _298

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive__(root *TreeNode) (result int) {

	var dfs func(node *TreeNode, prefix int, length int)
	dfs = func(node *TreeNode, prefix int, length int) {
		if node == nil {
			return
		}
		if node.Val == prefix+1 {
			length += 1
			result = max(result, length)
		} else {
			length = 1
		}

		dfs(node.Left, node.Val, length)
		dfs(node.Right, node.Val, length)
	}

	result = 1
	dfs(root.Left, root.Val, 1)
	dfs(root.Right, root.Val, 1)
	return
}
