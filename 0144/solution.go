package _144

func _preorderTraversal(root *TreeNode) (vals []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			vals = append(vals, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return
}
