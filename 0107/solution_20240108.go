package _107

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {

	results := make([][]int, 0)

	if root == nil {
		return results
	}

	nexts := make([]*TreeNode, 0)
	nexts = append(nexts, root)
	for len(nexts) != 0 {
		currents := nexts
		nexts = nexts[0:0:0]
		result := make([]int, 0)
		for _, c := range currents {
			result = append(result, c.Val)
			if c.Left != nil {
				nexts = append(nexts, c.Left)
			}
			if c.Right != nil {
				nexts = append(nexts, c.Right)
			}
		}

		results = append(results, result)
	}

	finals := make([][]int, len(results))
	for i := range results {
		finals[len(finals)-1-i] = results[i]
	}

	return finals
}
