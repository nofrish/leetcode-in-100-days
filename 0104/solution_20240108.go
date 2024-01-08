package _104

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	ldepth := maxDepth(root.Left)
	rdepth := maxDepth(root.Right)

	if ldepth >= rdepth {
		return ldepth + 1
	}
	return rdepth + 1
}
