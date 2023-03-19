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
	return nil
}

func (l *List) Delete(index int) error {
	return nil
}

func (l *List) DeleteAll(element rune) {
}

func (l *List) Get(index int) (rune, error) {
	if index < 0 || index >= l.Length() {
		return '\x00', fmt.Errorf("Invalid index %d", index)
	}
	node := l.head
	pos := 0
	for pos != index {
		pos++
		node = node.next
	}
	return node.key, nil
}

func (l *List) Clone(index int) List {
	return List{}
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
