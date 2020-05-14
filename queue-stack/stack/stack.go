package stack

// Stack is a LIFO Data Structure.
type Stack struct {
	Data []interface{}
}

// Push insert an element into the stack.
func (s *Stack) Push(val interface{}) {
	s.Data = append(s.Data, val)
}

// Pop delete an element from the stack.
// Return true if the operation is successful.
func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}

	s.Data = s.Data[:len(s.Data)-1]
	return true
}

// Top get the top item from the stack.
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Data[len(s.Data)-1]
}

// IsEmpty checks whether the stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
