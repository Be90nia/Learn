package main

import "fmt"

type Node struct {
	value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{value: value}
}

//struct test
func TreeNodeTest1() {
	var root Node
	root = Node{value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	fmt.Println(root)
	root.Print()
	root.SetValue(100)
	root.Print()

	root.Traverse()
}

//值接收者,拷贝值并不会修改参数的值
func (node Node) Print() {
	fmt.Print(node.value, "\n")
}

//指针接收者,引用参数值并且修改参数值
func (node *Node) SetValue(value int) {
	node.value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
