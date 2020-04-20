package ll3

// use doubly linked list
type MyLinkedList struct {
	len  int
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
	prev *Node
}

// node can not be head in linked list
// because we can not change head of
// linked in this func
func (p *Node) insertPrev(val int) *Node {
	np := &Node{
		data: val,
		prev: p.prev,
		next: p,
	}
	if p.prev != nil {
		p.prev.next = np
	}
	p.prev = np

	return np
}

// node can not be tail in linked list
// because we can not change tail of
// linked in this func
func (p *Node) insertNext(val int) *Node {
	np := &Node{
		data: val,
		prev: p,
		next: p.next,
	}
	if p.next != nil {
		p.next.prev = np
	}
	p.next = np

	return np
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.len {
		return -1
	}

	if index == 0 {
		return l.head.data
	}

	if index == l.len-1 {
		return l.tail.data
	}

	return l.find(index).data
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	if l.head == nil {
		l.head = &Node{data: val}
		l.tail = l.head
	} else {
		l.head = l.head.insertPrev(val)
	}

	l.len++
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	if l.tail == nil {
		l.tail = &Node{data: val}
		l.head = l.tail
	} else {
		l.tail = l.tail.insertNext(val)
	}

	l.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	// fix index less than 0
	if index < 0 {
		index = 0
	}

	// index can equal to length
	if index < 0 || index > l.len {
		return
	}

	if index == 0 {
		l.AddAtHead(val)
		return
	}

	if index == l.len {
		l.AddAtTail(val)
		return
	}

	l.find(index).insertPrev(val)
	l.len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.len {
		return
	}

	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
	} else if index == l.len-1 {
		l.tail = l.tail.prev
		if l.tail != nil {
			l.tail.next = nil
		}
	} else {
		p := l.find(index)
		// p must not be nil, so we don't need to judge
		p.next.prev = p.prev
		p.prev.next = p.next
	}

	l.len--
}

// note: return node maybe nil
func (l *MyLinkedList) find(index int) (p *Node) {
	var i int
	if index <= l.len-index {
		// start at head
		for i, p = 0, l.head; i < index; i, p = i+1, p.next {
		}
	} else {
		// start at tail
		for i, p = l.len-1, l.tail; i > index; i, p = i-1, p.prev {
		}
	}
	return
}
