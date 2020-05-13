package queue

// Queue is a FIFO Data Structure.
type Queue struct {
	data []interface{}
}

// EnQueue insert an element into the queue.
// Return nil true if the operation is successful.
func (q *Queue) EnQueue(val interface{}) bool {
	q.data = append(q.data, val)
	return true
}

// DeQueue delete an element from the queue.
// Return nil true if the operation is successful.
func (q *Queue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.data = q.data[1:]
	return true
}

// Front get the front item from the queue.
// If the queue is empty, return nil.
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data[0]
}

// IsEmpty checks whether the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
