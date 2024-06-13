package linkedlist

type node[T any] struct {
	value T
	next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
		next:  nil,
	}
}

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func New[T any](val T) *LinkedList[T] {
	node := newNode(val)

	return &LinkedList[T]{
		head:   node,
		tail:   node,
		length: 1,
	}
}

func (l *LinkedList[T]) Push(val T) *LinkedList[T] {
	newNode := &node[T]{value: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++
	return l
}

func (l *LinkedList[T]) Print(itr func(val T)) {
	temp := l.head

	if temp == nil {
		return
	}

	for temp.next != nil {
		itr(temp.value)
		temp = temp.next
	}
	itr(temp.value)
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Pop() *node[T] {

	if l.length == 0 {
		return nil
	}
	temp := l.head
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return temp
	}
	var prev *node[T] = nil

	for temp.next != nil {
		prev = temp
		temp = temp.next
	}
	deletedNode := temp
	prev.next = nil
	l.length--

	return deletedNode

}

func (l *LinkedList[T]) Get(index int) *node[T] {
	if index > l.length || index < 0 {
		return nil
	} else {
		temp := l.head
		for i := 0; i < index; i++ {
			temp = temp.next
		}
		return temp
	}

}

func (l *LinkedList[T]) GetValueAtAnyIndex(index int) T {
	var null T
	if index > l.length || index < 0 {
		return null
	} else {
		temp := l.head
		for i := 0; i < index; i++ {
			temp = temp.next
		}
		return temp.value
	}

}

func (l *LinkedList[T]) Unshift(value T) {
	node := newNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node

	}
	l.length++
}

func (l *LinkedList[T]) Shift() T {
	var val T
	if l.head != nil {
		temp := l.head
		l.head = l.head.next
		l.length--
		return temp.value
	}
	return val
}

func (l *LinkedList[T]) Insert(index int, value T) {
	if index > l.length || index < 0 {
		return
	} else if l.length == 0 || l.head == nil {
		l.Unshift(value)

	} else if l.length == index {
		l.Push(value)

	} else {
		node := l.Get(index - 1)
		n := newNode(value)
		n.next = node.next
		node.next = n

	}

	l.length++

}

func (l *LinkedList[T]) Reverse() {
	var prev *node[T]
	current := l.head
	var next *node[T]

	l.tail = l.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func (l *LinkedList[T]) GetMiddleNodeValue() T {

	var zeroValue T

	if l.head == nil {
		return zeroValue
	}
	if l.length == 1 {
		return l.head.value
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.value

}

func (l *LinkedList[T]) GetNthFromLast(n int) T {
	temp := l.head
	var hare *node[T]

	i := 1

	for temp != nil {
		if i == n {
			hare = temp
			i = 0
		}
		temp = temp.next
		i += 1
	}
	var zeroVal T

	if hare == nil {
		return zeroVal
	}

	return hare.value

}
