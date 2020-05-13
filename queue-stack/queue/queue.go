package queue

// Queue is a FIFO Data Structure.
type Queue struct {
	Data []interface{}
}

// EnQueue insert an element into the queue.
// Return nil true if the operation is successful.
func (q *Queue) EnQueue(val interface{}) bool {
	q.Data = append(q.Data, val)
	return true
}

// DeQueue delete an element from the queue.
// Return nil true if the operation is successful.
func (q *Queue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.Data = q.Data[1:]
	return true
}

// Front get the front item from the queue.
// If the queue is empty, return nil.
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Data[0]
}

// IsEmpty checks whether the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(q.Data) == 0
}
