package _226

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

func invertTree(root *TreeNode) *TreeNode {
	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		nodes := nexts
		nexts = make([]*TreeNode, 0)
		for _, node := range nodes {
			if node != nil {
				node.Left, node.Right = node.Right, node.Left
				nexts = append(nexts, node.Left)
				nexts = append(nexts, node.Right)
			}
		}
	}
	return root
}
