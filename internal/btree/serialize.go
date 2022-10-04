package btree

import (
	"encoding/json"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type Codec struct {
//	PreOrder []int `json:"pre_order"`
//	InOrder  []int `json:"in_order"`
//}
//
//func Constructor() Codec {
//	return Codec{}
//}

// Serializes a tree to a single string. 无重复元素情况下
//func (c *Codec) serialize(root *TreeNode) string {
//	var (
//		preOrderTraversal func(node *TreeNode)
//		inOrderTraversal  func(node *TreeNode)
//	)
//	preOrderTraversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		c.PreOrder = append(c.PreOrder, node.Val)
//		preOrderTraversal(node.Left)
//		preOrderTraversal(node.Right)
//	}
//	inOrderTraversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		inOrderTraversal(node.Left)
//		c.InOrder = append(c.InOrder, node.Val)
//		inOrderTraversal(node.Right)
//	}
//	preOrderTraversal(root)
//	inOrderTraversal(root)
//	return c.string()
//}

// Deserializes your encoded data to tree. 无重复元素情况下
//func (c *Codec) deserialize(data string) *TreeNode {
//	_ = json.Unmarshal([]byte(data), c)
//
//	var rebuildTree func(preOrder, inOrder []int) *TreeNode
//	rebuildTree = func(preOrder, inOrder []int) *TreeNode {
//		if len(preOrder) == 0 {
//			return nil
//		}
//		node := &TreeNode{}
//		node.Val = preOrder[0]
//		i := findIndex(inOrder, node.Val)
//		preOrder = preOrder[1:]
//		node.Left = rebuildTree(preOrder[:i], inOrder[:i])
//		node.Right = rebuildTree(preOrder[i:], inOrder[i+1:])
//		return node
//	}
//	return rebuildTree(c.PreOrder, c.InOrder)
//}

//func (c *Codec) string() string {
//	bytesData, _ := json.Marshal(c)
//	return string(bytesData)
//}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	// ["1","2","null","3","5"]
	var vals []string
	queue := []*TreeNode{root}
	for len(queue) > 0 && !c.allNil(queue) {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node == nil {
				vals = append(vals, "null")
				queue = append(queue, nil)
				queue = append(queue, nil)
			} else {
				vals = append(vals, strconv.Itoa(node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return c.string(vals)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	var vals []string
	_ = json.Unmarshal([]byte(data), &vals)

	var buildTree func(i int) *TreeNode
	buildTree = func(i int) *TreeNode {
		if i > len(vals) || vals[i-1] == "null" {
			return nil
		}
		node := &TreeNode{}
		node.Val, _ = strconv.Atoi(vals[i-1])
		node.Left = buildTree(2 * i)
		node.Right = buildTree(2*i + 1)
		return node
	}
	return buildTree(1)
}

func (c *Codec) string(vals []string) string {
	data, _ := json.Marshal(vals)
	return string(data)
}

func (c *Codec) trimNil(nodes []*TreeNode) []*TreeNode {
	if len(nodes) == 0 {
		return nodes
	}
	for nodes[len(nodes)-1] == nil {
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func (c *Codec) allNil(nodes []*TreeNode) bool {
	for _, node := range nodes {
		if node != nil {
			return false
		}
	}
	return true
}
