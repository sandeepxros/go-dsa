package main

import (
	"github.com/sandeepxros/go-dsa/queue"
)

func main() {
	var queue queue.Queue[int]

	for i := 0; i < 10; i++ {
		queue.Enque(i)
	}

	queue.ForEach(func(data int) {
		println(data)
	})

	println("deque-----------------------------___________>>>>>>>>>>>>>>>")
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Dequeue())
	println(queue.Peek())

	println("quwuw-----------------------------___________>>>>>>>>>>>>>>>")

	queue.ForEach(func(data int) {
		println(data)
	})

}
