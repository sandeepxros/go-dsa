package stack

import (
	linkedlist "github.com/sandeepxros/go-dsa/linkedList"
)

type Stack[T any] struct {
	list *linkedlist.LinkedList[T]
}

func (s *Stack[T]) Push(val T) {
	if s.list == nil {
		s.list = linkedlist.New(val)
	} else {
		s.list.Unshift(val)
	}
}

func (s *Stack[T]) Pop() T {
	return s.list.Shift()
}

func (s *Stack[T]) Len() int {
	if s.list == nil {
		return 0
	}
	return s.list.Len()
}

func (s *Stack[T]) Peek() T {
	var null T
	if s.list.Len() <= 0 {
		return null
	}
	return s.list.GetValueAtAnyIndex(0)
}

func (s *Stack[T]) ForEach(itr func(data T)) {
	s.list.Print(itr)
}
