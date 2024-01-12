package _448

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

func goodNodes(root *TreeNode) int {
	return 1 + traverse(root.Left, root.Val) + traverse(root.Right, root.Val)
}

func traverse(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

	good := 0
	if node.Val >= max {
		max = node.Val
		good = 1
	}
	return good + traverse(node.Left, max) + traverse(node.Right, max)
}
