package main

import "github.com/sandeepxros/go-dsa/doublyLinkedList"

func main() {
	dbList := doublyLinkedList.New(44)
	for i := 0; i < 10; i++ {
		dbList.Push(i)
	}
	dbList.Print(func(val int) {
		print(val, "     ")
	})
	println()
	dbList.PrintReverse(func(val int) {
		print(val, "     ")
	})

	for dbList.Len() > 0 {
		dbList.Pop()
	}

	dbList.PrintReverse(func(val int) {
		print(val, "     ")
	})

	println("list len", dbList.Len())
}
