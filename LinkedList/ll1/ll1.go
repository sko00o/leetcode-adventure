package ll1

// use slice
type MyLinkedList struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		data: make([]int, 0),
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(l.data) {
		return -1
	}
	return l.data[index]
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	l.data = append([]int{val}, l.data...)
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	l.data = append(l.data, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	// fix index less than 0
	if index < 0 {
		index = 0
	}

	// index can equal to length
	if index < 0 || index > len(l.data) {
		return
	}

	tmp := make([]int, index)
	copy(tmp, l.data[:index])
	l.data = append(append(tmp, val), l.data[index:]...)
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= len(l.data) {
		return
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
}
