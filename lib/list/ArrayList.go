package list

type AList[T any] struct {
	maxSize   int
	size      int
	curr      int
	listArray []T
}

func (l *AList[T]) Clear() {
	l.curr = 0
	l.size = 0
	l.listArray = make([]T, l.maxSize)
}

func (l *AList[T]) Insert(it T) {
	for i := l.size; i > l.curr; i-- {
		l.listArray[i] = l.listArray[i-1]
	}

	l.listArray[l.curr] = it
	l.size++
}

func (l *AList[T]) Append(it T) {
	l.size++
	l.listArray[l.size] = it
}

func (l *AList[T]) Remove() T {
	it := l.listArray[l.curr]

	for i := l.curr; i < l.size-1; i++ {
		l.listArray[i] = l.listArray[i+1]
	}

	l.size--

	return it
}

func (l AList[T]) Length() int {
	return l.size
}

func (l *AList[T]) Prev() {
	if l.curr != 0 {
		l.curr--
	}
}

func (l *AList[T]) Next() {
	if l.curr < l.size {
		l.curr++
	}
}

func (l AList[T]) CurrPos() int {
	return l.curr
}

func (l *AList[T]) MoveToStart() {
	l.curr = 0
}

func (l *AList[T]) MoveToEnd() {
	l.curr = l.size
}

func (l *AList[T]) MoveToPos(pos int) {
	l.curr = pos
}

func (l AList[T]) GetValue() T {
	return l.listArray[l.curr]
}
