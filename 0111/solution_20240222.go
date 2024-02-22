package _111

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := minDepth(root.Left)
	d2 := minDepth(root.Right)
	if d1 == 0 || d2 == 0 {
		return max(d1, d2) + 1
	}
	return min(d1, d2) + 1
}

func minDepth_(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nexts := []*TreeNode{root}
	for level := 1; len(nexts) > 0; level++ {
		nodes := nexts
		nexts = make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				nexts = append(nexts, node.Left)
			}
			if node.Right != nil {
				nexts = append(nexts, node.Right)
			}
		}
	}
	return -1
}
