package queue

type Link struct {
	element int
	next    *Link
}

type LQueue struct {
	size  int
	front *Link
	rear  *Link
}

func NewLQueue() *LQueue {
	headerNode := new(Link) // Use a header node to ease negation logic
	front := headerNode
	rear := front

	return &LQueue{
		front: front,
		rear:  rear,
	}
}

func (q *LQueue) Clear() {
	curr := q.front.next

	for curr != nil {
		temp := curr
		curr = curr.next
		temp.next = nil
	}

	q.front.next = curr
	q.rear = q.front
	q.size = 0
}

func (q *LQueue) Enqueue(it int) {
	q.rear.next = &Link{element: it, next: nil}
	q.rear = q.rear.next
	q.size++
}

func (q *LQueue) Dequeue() int {
	// return an err if the queue is empty
	temp := q.front.next
	q.front.next = q.front.next.next

	if q.rear == temp { // if temp was the last node
		q.rear = q.front
	}

	q.size--

	return temp.element
}

func (q LQueue) FrontValue() int {
	// return an err if the queue is empty
	return q.front.element
}

func (q LQueue) Length() int {
	return q.size
}
