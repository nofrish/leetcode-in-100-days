package _513

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {

	result := root.Val

	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		result = nexts[0].Val
		nodes := nexts
		nexts = nexts[0:0:0]
		for _, node := range nodes {
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
