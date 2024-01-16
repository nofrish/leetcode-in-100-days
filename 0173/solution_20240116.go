package _173

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

type BSTIterator struct {
	list []int
	cur  int
}

func Constructor(root *TreeNode) BSTIterator {

	iter := BSTIterator{
		list: make([]int, 0),
		cur:  0,
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}
		iter.list = append(iter.list, node.Val)
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)

	return iter
}

func (this *BSTIterator) Next() int {
	result := this.list[this.cur]
	this.cur++
	return result
}

func (this *BSTIterator) HasNext() bool {
	return this.cur < len(this.list)
}
