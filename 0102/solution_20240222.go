package _102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder_(root *TreeNode) (result [][]int) {
	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		nodes := nexts
		nexts = make([]*TreeNode, 0)

		vals := make([]int, 0)
		for _, node := range nodes {
			if node != nil {
				vals = append(vals, node.Val)
				nexts = append(nexts, node.Left)
				nexts = append(nexts, node.Right)
			}
		}
		if len(vals) > 0 {
			result = append(result, vals)
		}
	}
	return
}
