package _257

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths__(root *TreeNode) (result []string) {

	var dfs func(node *TreeNode, prefix string)
	dfs = func(node *TreeNode, prefix string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			result = append(result, fmt.Sprintf("%s%d", prefix, node.Val))
			return
		}
		dfs(node.Left, fmt.Sprintf("%s%d->", prefix, node.Val))
		dfs(node.Right, fmt.Sprintf("%s%d->", prefix, node.Val))
	}

	dfs(root, "")
	return
}
