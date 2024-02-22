package _107

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom_(root *TreeNode) (result [][]int) {

	if root == nil {
		return
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
		result = append(result, vals)
	}

	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return
}
