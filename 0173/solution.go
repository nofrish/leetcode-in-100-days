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
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{make([]*TreeNode, 0, 100)}
	iter.pushAllLeft(root)
	return iter
}

func (this *BSTIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.pushAllLeft(top.Right)
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) pushAllLeft(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
