package list

type Link[T any] struct {
	element T
	next    *Link[T]
}

type LList[T any] struct {
	head  *Link[T]
	tail  *Link[T]
	curr  *Link[T]
	count int
}

func (l LList[T]) GetValue() T {
	return l.curr.next.element
}

func (l LList[T]) Length() int {
	return l.count
}

func (l *LList[T]) Insert(it T) {
	l.curr.next = &Link[T]{element: it, next: l.curr.next}

	if l.tail == l.curr {
		l.tail = l.curr.next
	}

	l.count++
}

func (l *LList[T]) Append(it T) {
	l.tail.next = &Link[T]{element: it, next: l.curr.next}
	l.tail = l.tail.next
	l.count++
}

func (l *LList[T]) Remove() T {
	it := l.curr.next.element

	if l.tail == l.curr.next {
		l.tail = l.curr
	}

	l.curr.next = l.curr.next.next
	l.count--

	return it
}

func (l LList[T]) CurrPos() int {
	temp := l.head
	var i int

	for i = 0; l.curr != temp; i++ {
		temp = temp.next
	}

	return i
}

func (l *LList[T]) Prev() {
	if l.curr == l.head {
		return
	}

	temp := l.head

	for temp.next != l.curr {
		temp = temp.next
	}

	l.curr = temp
}

func (l *LList[T]) Next() {
	if l.curr != l.tail {
		l.curr = l.curr.next
	}
}

func (l *LList[T]) MoveToPos(pos int) {
	l.curr = l.head

	for i := 0; i < pos; i++ {
		l.curr = l.curr.next
	}
}

func (l *LList[T]) MoveToStart() {
	l.curr = l.head
}

func (l *LList[T]) MoveToEnd() {
	l.curr = l.tail
}

func (l *LList[T]) Clear() {
	for l.head != nil {
		l.head = l.head.next
	}

	x := new(Link[T])
	l.curr = x
	l.tail = x
	l.head = x
	l.count = 0
}

func (l LList[T]) Print() {}
