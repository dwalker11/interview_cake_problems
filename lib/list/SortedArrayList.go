package list

type SAList[T comparable] struct {
	maxSize   int
	size      int
	curr      int
	listArray []T
}

func (l *SAList[T]) Clear() {
	l.curr = 0
	l.size = 0
	l.listArray = make([]T, l.maxSize)
}

func (l *SAList[T]) Insert(it T) {
	for l.MoveToStart(); l.CurrPos() < l.Length(); l.Next() {
		curr := l.GetValue()
		if curr > it {
			break
		}
	}

	for i := l.size; i > l.curr; i-- {
		l.listArray[i] = l.listArray[i-1]
	}

	l.listArray[l.curr] = it
	l.size++
}

func (l *SAList[T]) Append(it T) {
	l.size++
	l.listArray[l.size] = it
}

func (l *SAList[T]) Remove() T {
	it := l.listArray[l.curr]

	for i := l.curr; i < l.size-1; i++ {
		l.listArray[i] = l.listArray[i+1]
	}

	l.size--

	return it
}

func (l SAList[T]) Length() int {
	return l.size
}

func (l *SAList[T]) Prev() {
	if l.curr != 0 {
		l.curr--
	}
}

func (l *SAList[T]) Next() {
	if l.curr < l.size {
		l.curr++
	}
}

func (l SAList[T]) CurrPos() int {
	return l.curr
}

func (l *SAList[T]) MoveToStart() {
	l.curr = 0
}

func (l *SAList[T]) MoveToEnd() {
	l.curr = l.size
}

func (l *SAList[T]) MoveToPos(pos int) {
	l.curr = pos
}

func (l SAList[T]) GetValue() T {
	return l.listArray[l.curr]
}
