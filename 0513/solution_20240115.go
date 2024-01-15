package _513

func findBottomLeftValue_(root *TreeNode) int {

	nexts := []*TreeNode{root}
	result := 0

	for len(nexts) > 0 {
		nodes := nexts
		result = nodes[0].Val
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
