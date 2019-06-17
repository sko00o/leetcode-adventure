package ll2

// use single linked list
type MyLinkedList struct {
	len  int
	head *Node
}

type Node struct {
	data int
	next *Node
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

	for i, p := 0, l.head; i <= index && p != nil; i, p = i+1, p.next {
		if i == index {
			return p.data
		}
	}

	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	p := &Node{
		data: val,
		next: l.head,
	}
	l.head = p
	l.len++
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.head = &Node{
			data: val,
		}
		l.len++
		return
	}

	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{
		data: val,
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

	for i, p, pre := 0, l.head, l.head; i <= index; i, p = i+1, p.next {
		if i == index {
			pre.next = &Node{
				data: val,
				next: p,
			}
			l.len++
			return
		}
		pre = p
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.len {
		return
	}

	if index == 0 {
		l.head = l.head.next
		l.len--
		return
	}

	for i, p, pre := 0, l.head, l.head; i <= index; i, p = i+1, p.next {
		if i == index {
			pre.next = p.next
			l.len--
			return
		}
		pre = p
	}
}
