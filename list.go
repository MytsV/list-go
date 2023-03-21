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
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *List) Insert(element rune, index int) error {
	if index < 0 || index > l.Length() {
		return fmt.Errorf("Invalid index %d", index)
	}
	if index == l.Length() {
		l.Append(element)
		return nil
	}

	inserted := &Node{key: element}
	current := l.getNode(index)
	inserted.next = current

	if index == 0 {
		l.head = inserted
	} else {
		current.prev.next = inserted
		inserted.prev = current.prev
		current.prev = inserted
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
		deleted := l.getNode(index)
		deleted.prev.next = deleted.next
		deleted.next.prev = deleted.prev
	}
	return nil
}

func (l *List) DeleteAll(element rune) {
	node := l.head
	pos := 0
	for pos < l.Length() {
		if node.key == element {
			node = node.next
			err := l.Delete(pos)
			if err != nil {
				panic("Can't delete all elements because of internal error")
			}
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
	clone := &List{}
	node := l.head
	for node != nil {
		clone.Append(node.key)
		node = node.next
	}
	return clone
}

func (l *List) Reverse() {
	var temp *Node
	current := l.head

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		l.head = temp.prev
	}
}

func (l *List) FindFirst(element rune) int {
	node := l.head
	pos := 0
	for node != nil {
		if node.key == element {
			return pos
		}
		pos++
		node = node.next
	}
	return -1
}

func (l *List) FindLast(element rune) int {
	node := l.tail
	pos := l.Length() - 1
	for node != nil {
		if node.key == element {
			return pos
		}
		pos--
		node = node.prev
	}
	return -1
}

func (l *List) Clear() {
	l.head = nil
	l.tail = nil
}

func (l *List) Extend(elements List) {
	node := elements.head
	for node != nil {
		l.Append(node.key)
		node = node.next
	}
}

// "Private" method. Isn't included in the interface
func (l *List) getNode(index int) *Node {
	node := l.head
	pos := 0
	for pos != index {
		pos++
		node = node.next
	}
	return node
}
