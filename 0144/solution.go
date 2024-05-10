package _144

// Solution One. Recursively.
func _preorderTraversal(root *TreeNode) (vals []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// Solution Two. Iteratively.
func __preorderTraversal(root *TreeNode) (vals []int) {
	nexts := []*TreeNode{root}
	for len(nexts) > 0 {
		node := nexts[len(nexts)-1]
		nexts = nexts[:len(nexts)-1]
		if node != nil {
			vals = append(vals, node.Val)
			nexts = append(nexts, node.Right)
			nexts = append(nexts, node.Left)
		}
	}
	return
}
