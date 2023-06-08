package list

type DLink[T any] struct {
	element T
	next    *DLink[T]
	prev    *DLink[T]
}

type DLList[T any] struct {
	head  *DLink[T]
	tail  *DLink[T]
	curr  *DLink[T]
	count int
}

func (l DLList[T]) GetValue() T {
	return l.curr.next.element
}

func (l DLList[T]) Length() int {
	return l.count
}

func (l *DLList[T]) Insert(it T) {
	l.curr.next.prev = &DLink[T]{element: it, prev: l.curr, next: l.curr.next}
	l.curr.next = l.curr.next.prev
	l.count++
}

func (l *DLList[T]) Append(it T) {
	l.tail.prev.next = &DLink[T]{element: it, next: l.curr.next}
	l.tail.prev = l.tail.prev.next
	l.count++
}

func (l *DLList[T]) Remove() (T, bool) {
	if l.curr.next == l.tail {
		var zero T
		return zero, false
	}

	it := l.curr.next.element
	l.curr.next.next.prev = l.curr
	l.curr.next = l.curr.next.next
	l.count--

	return it, true
}

func (l DLList[T]) CurrPos() int {
	temp := l.head
	var i int

	for i = 0; l.curr != temp; i++ {
		temp = temp.next
	}

	return i
}

func (l *DLList[T]) Prev() {
	if l.curr != l.head {
		l.curr = l.curr.prev
	}
}

func (l *DLList[T]) Next() {
	if l.curr != l.tail {
		l.curr = l.curr.next
	}
}

func (l *DLList[T]) MoveToPos(pos int) {
	l.curr = l.head

	for i := 0; i < pos; i++ {
		l.curr = l.curr.next
	}
}

func (l *DLList[T]) MoveToStart() {
	l.curr = l.head
}

func (l *DLList[T]) MoveToEnd() {
	l.curr = l.tail
}

func (l *DLList[T]) Clear() {
	for l.head != nil {
		l.head = l.head.next
	}

	x := new(DLink[T])
	l.curr = x
	l.tail = x
	l.head = x
	l.count = 0
}

func (l DLList[T]) Print() {}
