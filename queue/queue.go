package queue

import "github.com/sandeepxros/go-dsa/doublyLinkedList"

type Queue[T any] struct {
	list *doublyLinkedList.LinkedList[T]
}

func (s *Queue[T]) Enque(val T) {
	if s.list == nil {
		s.list = doublyLinkedList.New(val)
	} else {
		s.list.Push(val)
	}
}

func (s *Queue[T]) Dequeue() T {
	return s.list.Shift()
}

func (s *Queue[T]) Len() int {
	if s.list == nil {
		return 0
	}
	return s.list.Len()
}

func (s *Queue[T]) Peek() T {
	var null T
	if s.list.Len() <= 0 {
		return null
	}
	return s.list.GetValueAtAnyIndex(0)
}

func (s *Queue[T]) ForEach(itr func(data T)) {
	s.list.Print(itr)
}
