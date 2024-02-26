package _108

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

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1 : len(nums)])

	return root
}
