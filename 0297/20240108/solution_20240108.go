package _0240108

import (
	"fmt"
	"strconv"
	"strings"
)

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "-"
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	root, _ := helper(strings.Split(data, ","))
	return root
}

func helper(queue []string) (*TreeNode, []string) {

	elem := queue[0]
	queue = queue[1:]

	if elem == "-" {
		return nil, queue
	}

	val, _ := strconv.Atoi(elem)
	left, queue := helper(queue)
	right, queue := helper(queue)

	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}, queue
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
