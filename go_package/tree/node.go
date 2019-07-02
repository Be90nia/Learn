package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

//值接收者,拷贝值并不会修改参数的值
func (node Node) Print() {
	fmt.Print(node.Value, "\n")
}

//指针接收者,引用参数值并且修改参数值
func (node *Node) SetValue(Value int) {
	node.Value = Value
}

func (node *Node) Traverse() {
	// if node == nil {
	// 	return
	// }

	// node.Left.Traverse()
	// node.Print()
	// node.Right.Traverse()
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node{
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
