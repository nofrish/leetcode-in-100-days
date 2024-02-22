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

func binaryTreePaths_(root *TreeNode) (result []string) {

	var dfs func(node *TreeNode, prefix string)
	dfs = func(node *TreeNode, prefix string) {
		prefix = fmt.Sprintf("%s%d->", prefix, node.Val)
		if node.Left == nil && node.Right == nil {
			result = append(result, prefix[:len(prefix)-2]) // remove trailing "->"
		}
		if node.Left != nil {
			dfs(node.Left, prefix)
		}
		if node.Right != nil {
			dfs(node.Right, prefix)
		}
	}
	dfs(root, "")
	return
}
