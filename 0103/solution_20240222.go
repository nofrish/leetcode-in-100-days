package _103

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

func zigzagLevelOrder(root *TreeNode) (result [][]int) {

	if root == nil {
		return
	}

	flag := 1
	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		nodes := nexts
		nexts = make([]*TreeNode, 0)

		vals := make([]int, 0)
		for _, node := range nodes {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nexts = append(nexts, node.Left)
			}
			if node.Right != nil {
				nexts = append(nexts, node.Right)
			}
		}
		if flag < 0 {
			reverse(vals)
		}
		result = append(result, vals)
		flag = -flag
	}
	return
}
