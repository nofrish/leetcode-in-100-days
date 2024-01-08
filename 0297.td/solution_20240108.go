package _297_td

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
		return ""
	}

	result := fmt.Sprintf("%d,", root.Val)

	if root.Left != nil {
		result += fmt.Sprintf("%d,", root.Left.Val)
	} else {
		result += "-,"
	}

	if root.Right != nil {
		result += fmt.Sprintf("%d", root.Right.Val)
	} else {
		result += "-"
	}

	leftSeria := this.serialize(root.Left)
	rightSeria := this.serialize(root.Right)

	if leftSeria != "" {
		result = result + ";" + leftSeria
	}
	if rightSeria != "" {
		result = result + ";" + rightSeria
	}

	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if data == "" {
		return nil
	}
	nodeStrings := strings.Split(data, ";")

	rootString := nodeStrings[0]
	parts := strings.Split(rootString, ",")
	val, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: val}

	nexts := make([]*TreeNode, 0)
	if parts[1] != "-" {
		val, _ := strconv.Atoi(parts[1])
		root.Left = &TreeNode{Val: val}
		nexts = append(nexts, root.Left)
	}
	if parts[2] != "-" {
		val, _ := strconv.Atoi(parts[2])
		root.Right = &TreeNode{Val: val}
		nexts = append(nexts, root.Right)
	}

	for i := 1; len(nexts) > 0; i++ {
		node := nexts[0]
		nexts = nexts[1:]

		nodeString := nodeStrings[i]
		parts := strings.Split(nodeString, ",")
		if parts[1] != "-" {
			val, _ := strconv.Atoi(parts[1])
			node.Left = &TreeNode{Val: val}
			nexts = append(nexts, node.Left)
		}
		if parts[2] != "-" {
			val, _ := strconv.Atoi(parts[2])
			node.Right = &TreeNode{Val: val}
			nexts = append(nexts, node.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
