package _101

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric_(root *TreeNode) bool {

	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}

	return compare(root.Left, root.Right)
}
