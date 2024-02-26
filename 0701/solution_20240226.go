package _701

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func insertIntoBST_(root *TreeNode, val int) *TreeNode {

	node := &TreeNode{val, nil, nil}
	if root == nil {
		return node
	}

	var prev *TreeNode
	for cur := root; cur != nil; {
		prev = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if prev.Val > val {
		prev.Left = &TreeNode{val, nil, nil}
	} else {
		prev.Right = &TreeNode{val, nil, nil}
	}

	return root
}
