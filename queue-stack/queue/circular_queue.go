package queue

// CircularQueue can reused the wasted storage.
// It's a more efficient way to implement a queue.
type CircularQueue interface {
	Queue

	// Get the last item from the queue.
	// If the queue is empty, return -1.
	Rear() int

	// Checks whether the circular queue is full or not.
	IsFull() bool
}

// MyCircularQueue is circular queue struct.
type MyCircularQueue struct {
	data []int
	head int
	tail int
}

// Constructor set the size of the queue to be k.
func Constructor(k int) *MyCircularQueue {
	return &MyCircularQueue{
		data: make([]int, k),
		head: -1,
		tail: -1,
	}
}

// EnQueue insert an element into the circular queue.
// Return true if the operation is successful.
func (q *MyCircularQueue) EnQueue(x int) bool {
	if q.IsFull() {
		return false
	}
	if q.IsEmpty() {
		q.head = 0
	}

	q.tail = (q.tail + 1) % len(q.data)
	q.data[q.tail] = x
	return true
}

// DeQueue delete an element from the circular queue.
// Return true if the operation is successful.
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head = -1
		q.tail = -1
	} else {
		q.head = (q.head + 1) % len(q.data)
	}
	return true
}

// Front get the front item from the queue.
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

// Rear get the last item from the queue.
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.tail]
}

// IsEmpty checks whether the circular queue is empty or not.
func (q *MyCircularQueue) IsEmpty() bool {
	return q.head == -1
}

// IsFull checks whether the circular queue is full or not.
func (q *MyCircularQueue) IsFull() bool {
	return (q.tail+1)%len(q.data) == q.head
}
