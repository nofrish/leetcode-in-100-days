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

	nexts := make([]*TreeNode, 0)
	nexts = append(nexts, root)

	for len(nexts) > 0 {
		s := make([]int, 0)
		for _, node := range nexts {
			s = append(s, node.Val)
		}
		result = append(result, s)

		curs := nexts
		nexts = nexts[0:0:0]

		for _, node := range curs {
			if node.Left != nil {
				nexts = append(nexts, node.Left)
			}
			if node.Right != nil {
				nexts = append(nexts, node.Right)
			}
		}
	}

	return result
}
