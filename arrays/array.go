package arrays

type Array[T any] struct {
	List     []T
	Size     int
	Capacity int
}

func (a *Array[T]) New(capacity int) *Array[T] {
	return &Array[T]{
		List:     make([]T, capacity),
		Size:     0,
		Capacity: capacity,
	}
}

func (a *Array[T]) Push(val T) {
	a.List = append(a.List, val)
	a.Size++
}

func (a *Array[T]) Pop() {
	a.List = a.List[:a.Size]
	a.Size--
}

func (a *Array[T]) Unshift(val T) {

	a.Push(val)

	for i := range a.List {
		print(i, "	")
	}

	for i := a.Size - 1; i > 1; i-- {
		temp := a.List[i]
		a.List[i] = a.List[i-1]
		a.List[i-1] = temp
	}
}
