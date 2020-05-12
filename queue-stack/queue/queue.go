package queue

// Queue is FIFO data struct.
type Queue interface {
	// Insert an element into the queue.
	// Return true if the operation is successful.
	enQueue(x int) bool

	// Delete an element from the queue.
	// Return true if the operation is successful.
	deQueue() bool

	// Get the front item from the queue.
	Front() int

	// Checks whether the queue is empty or not.
	isEmpty() bool
}

type myQueue struct {
	data []int
}

func (q *myQueue) enQueue(x int) bool {
	q.data = append(q.data, x)
	return true
}

func (q *myQueue) deQueue() bool {
	if q.isEmpty() {
		return false
	}

	q.data = q.data[1:]
	return true
}

func (q myQueue) Front() int {
	return q.data[0]
}

func (q myQueue) isEmpty() bool {
	return len(q.data) == 0
}
