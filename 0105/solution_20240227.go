package _105

import (
	"slices"
)

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

func buildTree(preorder []int, inorder []int) *TreeNode {

	var dfs func(preStart, inStart, inEnd int) *TreeNode
	dfs = func(preStart, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		val := preorder[preStart]
		node := &TreeNode{val, nil, nil}
		mid := inStart + slices.Index(inorder[inStart:inEnd+1], val)
		node.Left = dfs(preStart+1, inStart, mid-1)
		node.Right = dfs(preStart+mid-inStart+1, mid+1, inEnd)
		return node
	}

	return dfs(0, 0, len(inorder)-1)
}
