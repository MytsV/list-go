package list

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  rune
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Length() int {
	length := 0
	node := l.head
	for node != nil {
		node = node.next
		length++
	}
	return length
}

func (l *List) Append(element rune) {
	node := &Node{key: element}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
}

func (l *List) Insert(element rune, index int) error {
	if index < 0 || index > l.Length() {
		return fmt.Errorf("Invalid index %d", index)
	}
	if index == l.Length() {
		l.Append(element)
		return nil
	}

	newNode := &Node{key: element}
	if index == 0 {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else {
		node := l.getNode(index)
		newNode.next = node
		node.prev.next = newNode
		newNode.prev = node.prev
		node.prev = newNode
	}
	return nil
}

func (l *List) Delete(index int) error {
	if index < 0 || index >= l.Length() {
		return fmt.Errorf("Invalid index %d", index)
	}
	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
	} else if index == l.Length()-1 {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		toDelete := l.getNode(index)
		toDelete.prev.next = toDelete.next
		toDelete.next.prev = toDelete.prev
	}
	return nil
}

func (l *List) DeleteAll(element rune) {
	node := l.head
	pos := 0
	for pos < l.Length() {
		if node.key == element {
			node = node.next
			l.Delete(pos)
		} else {
			pos++
			node = node.next
		}
	}
}

func (l *List) Get(index int) (rune, error) {
	if index < 0 || index >= l.Length() {
		return '\x00', fmt.Errorf("Invalid index %d", index)
	}
	node := l.getNode(index)
	return node.key, nil
}

func (l *List) Clone() *List {
	return &List{}
}

func (l *List) Reverse() {
}

func (l *List) FindFirst(element rune) int {
	return -1
}

func (l *List) FindLast(element rune) int {
	return -1
}

func (l *List) Clear() {
}

func (l *List) Extend(elements List) {
}

func (l *List) getNode(index int) *Node {
	node := l.head
	pos := 0
	for pos != index {
		pos++
		node = node.next
	}
	return node
}
