package _102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)
	if root == nil {
		return result
	}

	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		nodes := nexts
		nexts = nexts[0:0:0]
		vals := make([]int, len(nodes))
		for i, node := range nodes {
			vals[i] = node.Val
			if node.Left != nil {
				nexts = append(nexts, node.Left)
			}
			if node.Right != nil {
				nexts = append(nexts, node.Right)
			}
		}
		result = append(result, vals)
	}

	return result
}
