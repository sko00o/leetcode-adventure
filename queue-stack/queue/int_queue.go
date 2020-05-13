package queue

// IntQueue only process integers.
type IntQueue interface {
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

type myIntQueue struct {
	Queue
}

func (q *myIntQueue) EnQueue(x int) bool {
	return q.Queue.EnQueue(x)
}

func (q myIntQueue) Front() int {
	val, ok := q.Queue.Front().(int)
	if !ok {
		return -1
	}
	return val
}
