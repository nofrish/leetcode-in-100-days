package solution_20240511

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

type ColumnNode struct {
	Column int
	Node   *TreeNode
}

func verticalOrder(root *TreeNode) (result [][]int) {

	low, high := 1, 0
	columns := make(map[int][]int)

	nexts := []*ColumnNode{&ColumnNode{0, root}}
	for len(nexts) > 0 {
		nodes := nexts
		nexts = make([]*ColumnNode, 0)
		for _, node := range nodes {
			if node.Node == nil {
				continue
			}
			low, high = min(low, node.Column), max(high, node.Column)
			column := columns[node.Column]
			if column == nil {
				column = make([]int, 0)
			}
			columns[node.Column] = append(column, node.Node.Val)
			nexts = append(nexts, &ColumnNode{node.Column - 1, node.Node.Left})
			nexts = append(nexts, &ColumnNode{node.Column + 1, node.Node.Right})
		}
	}

	for i := low; i <= high; i++ {
		result = append(result, columns[i])
	}
	return
}
