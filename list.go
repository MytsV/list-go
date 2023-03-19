package list

type List struct {
}

func (l *List) Length() int {
	return -1
}

func (l *List) Append(element rune) {
}

func (l *List) Insert(element rune, index int) {
}

func (l *List) Delete(index int) {
}

func (l *List) DeleteAll(element rune) {
}

func (l *List) Get(index int) (rune, error) {
	return '\x00', nil
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
