package main

import (
	"fmt"

	avltrees "github.com/sandeepxros/go-dsa/avl-trees"
)

func main() {
	avltrees.Root = avltrees.Insert(avltrees.Root, 3)
	avltrees.Root = avltrees.Insert(avltrees.Root, 9)

	avltrees.Root = avltrees.Insert(avltrees.Root, 4)
	avltrees.Print(avltrees.Root)

	fmt.Println()

	avltrees.InorderTravasal(avltrees.Root)

	fmt.Println()
}

//30 42 42 43 48 78 90 90 100
