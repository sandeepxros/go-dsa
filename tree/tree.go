package tree

import (
	"github.com/sandeepxros/go-dsa/queue"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
	q    queue.Queue[*TreeNode]
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		data:  val,
		left:  nil,
		right: nil,
	}
}

func (l *Tree) New(val int) {
	node := newTreeNode(val)
	l.root = node
	l.q.Enque(node)
}

func (l *Tree) Insert(val int) {
	node := newTreeNode(val)
	if l.root == nil {
		l.root = node
		l.q.Enque(node)
		return
	}
	nodeFromQ := l.q.Peek()

	if nodeFromQ.left == nil {
		nodeFromQ.left = node
	} else {
		nodeFromQ.right = node
		l.q.Dequeue()
	}

	l.q.Enque(node)
}

func (l *Tree) Travarse(order int) {
	if order == 1 {
		preOrderTravarsal(l.root)
	} else if order == 2 {
		inorderTravarsal(l.root)
	} else if order == 3 {
		postOrderTravarsal(l.root)
	}
	println()
}

func preOrderTravarsal(t *TreeNode) {
	if t != nil {
		print(t.data, "	")
		preOrderTravarsal(t.left)
		preOrderTravarsal(t.right)
	}

}
func inorderTravarsal(t *TreeNode) {
	if t != nil {
		inorderTravarsal(t.left)
		print(t.data, "	")
		inorderTravarsal(t.right)
	}
}
func postOrderTravarsal(t *TreeNode) {
	if t != nil {
		postOrderTravarsal(t.left)
		postOrderTravarsal(t.right)
		print(t.data, "	")
	}
}
