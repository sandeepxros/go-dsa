package main

import (
	"github.com/sandeepxros/go-dsa/stack"
)

func main() {
	var st stack.Stack[int]
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	st.ForEach(func(data int) {
		println(data)
	})
}
