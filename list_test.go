package list

import (
	"testing"
)

func TestLength(t *testing.T) {
	list := List{}

	if list.Length() != 0 {
		t.Fatalf("Empty list's length does not equal 0")
	}

	list.Append('a')
	if list.Length() != 1 {
		t.Fatalf("List's length doesn't equal the count of new elements")
	}

	elementCount := 1024
	for i := 0; i < elementCount; i++ {
		list.Append('b')
	}
	if list.Length() != elementCount+1 {
		t.Fatalf("List has low limit to the count of elements")
	}
}

func TestAppend(t *testing.T) {
	list := List{}

	list.Append('a')
	if list.Length() != 1 {
		t.Fatalf("List's length doesn't change after appending")
	}
	value, err := list.Get(0)
	if err != nil {
		t.Fatalf("Character wasn't actually appended")
	} else if value != 'a' {
		t.Fatalf("Appended and real first character don't match")
	}

	list.Append('b')
	value, err = list.Get(0)
	if value != 'a' || err != nil {
		t.Fatalf("Existing element doesn't stay the same after appending a new one")
	}
}

func TestGet(t *testing.T) {
	list := List{}

	list.Append('a')
	list.Append('b')
	list.Append('c')
	list.Append('d')
	list.Append('e')

	value, err := list.Get(3)
	if err != nil {
		t.Fatalf("Can't retrieve an existing element")
	} else if value != 'd' {
		t.Fatalf("Wrong element is being retrieved")
	}

	if list.Length() != 5 {
		t.Fatalf("Retrieval produces side effects")
	}

	_, err = list.Get(5)
	if err == nil {
		t.Fatalf("No error produced on passing an invalid index")
	}
	_, err = list.Get(-1)
	if err == nil {
		t.Fatalf("No error produced on passing an invalid index")
	}
}
