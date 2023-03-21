package list

import "fmt"

type List struct {
	slice []rune
}

func (l *List) Length() int {
	return len(l.slice)
}

func (l *List) Append(element rune) {
	l.slice = append(l.slice, element)
}

func (l *List) Insert(element rune, index int) error {
	if index < 0 || index > l.Length() {
		return fmt.Errorf("Invalid index %d", index)
	}
	if index == l.Length() {
		l.Append(element)
		return nil
	}

	l.slice = append(l.slice[:index+1], l.slice[index:]...)
	l.slice[index] = element
	return nil
}

func (l *List) Delete(index int) error {
	if index < 0 || index >= l.Length() {
		return fmt.Errorf("Invalid index %d", index)
	}

	l.slice = append(l.slice[:index], l.slice[index+1:]...)
	return nil
}

func (l *List) DeleteAll(element rune) {
	for pos := 0; pos < len(l.slice); {
		if l.slice[pos] == element {
			err := l.Delete(pos)
			if err != nil {
				panic("Can't delete all elements because of internal error")
			}
		} else {
			pos++
		}
	}
}

func (l *List) Get(index int) (rune, error) {
	if index < 0 || index >= l.Length() {
		return '\x00', fmt.Errorf("Invalid index %d", index)
	}
	return l.slice[index], nil
}

func (l *List) Clone() *List {
	clone := &List{}
	cloneSlice := make([]rune, len(l.slice))
	copy(cloneSlice, l.slice)
	clone.slice = cloneSlice
	return clone
}

func (l *List) Reverse() {
	for i, j := 0, len(l.slice)-1; i < j; i, j = i+1, j-1 {
		l.slice[i], l.slice[j] = l.slice[j], l.slice[i]
	}
}

func (l *List) FindFirst(element rune) int {
	for pos := 0; pos < len(l.slice); pos++ {
		if l.slice[pos] == element {
			return pos
		}
	}
	return -1
}

func (l *List) FindLast(element rune) int {
	for pos := len(l.slice) - 1; pos >= 0; pos-- {
		if l.slice[pos] == element {
			return pos
		}
	}
	return -1
}

func (l *List) Clear() {
	l.slice = []rune{}
}

func (l *List) Extend(elements List) {
	l.slice = append(l.slice, elements.slice...)
}
