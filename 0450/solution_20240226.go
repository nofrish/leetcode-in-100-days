package _450

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteNode_(root *TreeNode, key int) *TreeNode {

	findMin := func(root *TreeNode) int {
		for root.Left != nil {
			root = root.Left
		}
		return root.Val
	}

	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		root.Val = findMin(root.Right)
		root.Right = deleteNode_(root.Right, root.Val)
	} else {
		if key < root.Val {
			root.Left = deleteNode_(root.Left, key)
		} else {
			root.Right = deleteNode_(root.Right, key)
		}
	}

	return root
}
