package _101

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
