package avltrees

import "fmt"

type Node struct {
	LChild *Node
	Height int
	Data   int
	RChild *Node
}

var Root *Node = nil

func createNode(val int) *Node {
	return &Node{
		LChild: nil,
		RChild: nil,
		Data:   val,
		Height: 1,
	}
}

func NodeHeight(p *Node) int {
	var hl, hr int
	if p != nil && p.LChild != nil {
		hl = p.LChild.Height
	} else {
		hl = 0
	}

	if p != nil && p.RChild != nil {
		hr = p.RChild.Height
	} else {
		hr = 0
	}

	if hl < hr {
		return hr + 1
	} else {
		return hl + 1
	}
}

func BalanceFactor(p *Node) int {
	var hl, hr int
	if p != nil && p.LChild != nil {
		hl = p.LChild.Height
	} else {
		hl = 0
	}

	if p != nil && p.RChild != nil {
		hr = p.RChild.Height
	} else {
		hr = 0
	}
	return hl - hr
}

func LLRotation(p *Node) *Node {
	pl := p.LChild
	plr := pl.RChild

	pl.RChild = p
	p.LChild = plr

	p.Height = NodeHeight(p)
	pl.Height = NodeHeight(pl)

	if Root == p {
		Root = pl
	}

	return pl
}

func RRRotation(p *Node) *Node {
	pr := p.RChild
	prl := pr.LChild

	pr.LChild = p
	p.RChild = prl

	p.Height = NodeHeight(p)
	pr.Height = NodeHeight(pr)

	if Root == p {
		Root = pr
	}

	return pr
}

func LRRotation(p *Node) *Node {
	//single rotation
	pr := p.RChild
	prl := pr.LChild

	prl.RChild = p
	p.RChild = nil

	prl.LChild = pr
	pr.LChild = nil

	if Root == p {
		Root = prl
	}

	Print(Root)

	return LLRotation(p)
}

func Insert(node *Node, val int) *Node {
	newNode := createNode(val)

	if node == nil {
		return newNode
	}

	if val < node.Data {
		node.LChild = Insert(node.LChild, val)
	} else if val > node.Data {
		node.RChild = Insert(node.RChild, val)
	}

	node.Height = NodeHeight(node)

	if BalanceFactor(node) == 2 && BalanceFactor(node.LChild) == 1 {
		return LLRotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.RChild) == -1 {
		return RRRotation(node)
	} else if BalanceFactor(node) == -2 && BalanceFactor(node.RChild) == 1 {
		return LRRotation(node)
	}

	return node
}

func Print(node *Node) {
	if node != nil {
		fmt.Print(node.Data, "	")
		Print(node.LChild)
		Print(node.RChild)
	}

}

func InorderTravasal(node *Node) {
	if node != nil {
		InorderTravasal(node.LChild)
		fmt.Print(node.Data, " ")
		InorderTravasal(node.RChild)
	}
}
