package queue

// Queue is FIFO data struct.
type Queue interface {
	// Insert an element into the queue.
	// Return true if the operation is successful.
	EnQueue(x int) bool

	// Delete an element from the queue.
	// Return true if the operation is successful.
	DeQueue() bool

	// Get the front item from the queue. 
	// If the queue is empty, return -1.
	Front() int

	// Checks whether the queue is empty or not.
	IsEmpty() bool
}

type myQueue struct {
	data []int
}

func (q *myQueue) EnQueue(x int) bool {
	q.data = append(q.data, x)
	return true
}

func (q *myQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.data = q.data[1:]
	return true
}

func (q myQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[0]
}

func (q myQueue) IsEmpty() bool {
	return len(q.data) == 0
}
