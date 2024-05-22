package main

import linkedlist "github.com/sandeepxros/go-dsa/linkedList"

func main() {
	list := linkedlist.New(55)
	list.Pop()
	for i := 0; i < 10; i++ {
		list.Push(i)
	}
	println("list len", list.Len())

	list.Print(func(val int) {
		print(val, "   ")
	})

	println()
	println("middle node of this linked list", list.GetMiddleNodeValue())
}
