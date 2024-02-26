package _099

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

func recoverTree(root *TreeNode) {

	var prev, first, second *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)

			if prev != nil && prev.Val > node.Val {
				if first == nil {
					first, second = prev, node
				} else {
					second = node
					return
				}
			}
			prev = node

			dfs(node.Right)
		}
	}

	dfs(root)
	first.Val, second.Val = second.Val, first.Val
	return
}
