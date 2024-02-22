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

func inorderTraversal__(root *TreeNode) (result []int) {

	nexts := make([]*TreeNode, 0)
	for root != nil || len(nexts) != 0 {
		if root != nil {
			nexts = append(nexts, root)
			root = root.Left
		} else {
			node := nexts[len(nexts)-1]
			nexts = nexts[:len(nexts)-1]
			result = append(result, node.Val)
			root = node.Right
		}
	}
	return
}
