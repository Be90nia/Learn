package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

//struct test
func TreeNodeTest1() {
	var root TreeNode
	root = TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)
	fmt.Println(root)
	root.Print()
	root.SetValue(100)
	root.Print()

	root.Traverse()
}

//值接收者,拷贝值并不会修改参数的值
func (node TreeNode) Print() {
	fmt.Print(node.value, "\n")
}

//指针接收者,引用参数值并且修改参数值
func (node *TreeNode) SetValue(value int) {
	node.value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}
