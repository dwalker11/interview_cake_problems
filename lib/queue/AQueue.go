package queue

type AQueue struct {
	maxSize   int
	front     int
	rear      int
	listArray []int
}

func NewAQueue(size int) *AQueue {
	return &AQueue{
		maxSize:   size + 1,
		front:     1,
		rear:      0,
		listArray: make([]int, size+1),
	}
}

func (q *AQueue) Clear() {
	q.rear = 0
	q.front = 1
}

func (q *AQueue) Enqueue(val int) {
	// return an err if queue is full
	q.rear = (q.rear + 1) % q.maxSize
	q.listArray[q.rear] = val
}

func (q *AQueue) Dequeue() int {
	// return an err if queue is empty
	val := q.listArray[q.front]
	q.front = (q.front + 1) % q.maxSize
	return val
}

func (q AQueue) FrontValue() int {
	// return an err if queue is empty
	return q.listArray[q.front]
}

func (q AQueue) Length() int {
	return ((q.rear + q.maxSize) - q.front + 1) % q.maxSize
}
