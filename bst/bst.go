package bst

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func newNode(val int) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

func Create(val int) *Node {
	return newNode(val)
}

func (n *Node) Insert(val int) {
	temp := n
	node := newNode(val)

	if n == nil {

		n = node

		println(n)
		return
	}

	for {

		if val < temp.value {
			if temp.left == nil {
				temp.left = node
				return
			}
			temp = temp.left

		} else {
			if temp.right == nil {
				temp.right = node
				return
			}
			temp = temp.right
		}
	}

}

func Print(n *Node) {
	if n == nil {
		fmt.Println("nothing in the list")
		return
	}
	fmt.Print(n.value, "	")
	if n.left != nil {
		Print(n.left)
	}

	if n.right != nil {
		Print(n.right)
	}
}

func (n *Node) Find(val int) *Node {
	if n == nil {
		return nil
	}
	temp := n
	for temp != nil {
		if temp.value == val {
			return temp
		}

		if temp.value < val && temp.right != nil {
			temp = temp.right
		}
		if temp.value > val && temp.left != nil {
			temp = temp.left
		}
	}
	return nil
}
