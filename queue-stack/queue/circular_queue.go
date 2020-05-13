package queue

// CircularQueue can reused the wasted storage.
// It's a more efficient way to implement a queue.
type CircularQueue interface {
	Queue

	// Get the last item from the queue.
	// If the queue is empty, return -1.
	Rear() int

	// Checks whether the circular queue is full or not.
	isFull() bool
}

type myCircularQueue struct {
	data []int
	head int
	tail int
}

// MyCircularQueue is a constructor, set the size of the queue to be k.
func MyCircularQueue(k int) CircularQueue {
	return &myCircularQueue{
		data: make([]int, k),
		head: -1,
		tail: -1,
	}
}

func (q *myCircularQueue) enQueue(x int) bool {
	if q.isFull() {
		return false
	}
	if q.isEmpty() {
		q.head = 0
	}

	q.tail = (q.tail + 1) % len(q.data)
	q.data[q.tail] = x
	return true
}

func (q *myCircularQueue) deQueue() bool {
	if q.isEmpty() {
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

func (q myCircularQueue) Front() int {
	if q.isEmpty() {
		return -1
	}

	return q.data[q.head]
}

func (q myCircularQueue) Rear() int {
	if q.isEmpty() {
		return -1
	}

	return q.data[q.tail]
}

func (q myCircularQueue) isEmpty() bool {
	return q.head == -1
}

func (q myCircularQueue) isFull() bool {
	return (q.tail+1)%len(q.data) == q.head
}
