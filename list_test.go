package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var stubElements = []rune{'a', 'b', 'c', 'd', 'e'}

func getStubElement(idx int) rune {
	stubIdx := idx % len(stubElements)
	return stubElements[stubIdx]
}

func getList(length int) List {
	list := List{}
	for i := 0; i < length; i++ {
		list.Append(getStubElement(i))
	}
	return list
}

func constructList(slice []rune) List {
	list := List{}
	for i := 0; i < len(slice); i++ {
		list.Append(slice[i])
	}
	return list
}

func areSame(l List, slice []rune) bool {
	if l.Length() != len(slice) {
		return false
	}
	for i := 0; i < len(slice); i++ {
		value, err := l.Get(i)
		if value != slice[i] || err != nil {
			return false
		}
	}
	return true
}

func TestLength(t *testing.T) {
	type testCase struct {
		msg    string
		length int
	}

	testCases := []testCase{
		{
			msg:    "Returns 0 for empty list",
			length: 0,
		},
		{
			msg:    "Changes return when a new element is appended",
			length: 1,
		},
		{
			msg:    "Returns correct value for any append count",
			length: 1024,
		},
	}

	for _, test := range testCases {
		list := getList(test.length)
		actual := list.Length()
		assert.Equal(t, test.length, actual, test.msg)
	}
}

func TestAppend(t *testing.T) {
	msg := "Assigns elements and follows order"
	length := 10
	list := getList(length) //getList function uses list.Append()
	assert := assert.New(t)
	for i := 0; i < length; i++ {
		value, err := list.Get(i)
		assert.Nil(err, msg)
		assert.Equal(getStubElement(i), value, msg)
	}
}

func TestGet(t *testing.T) {
	length := 15
	list := getList(length)
	type testCase struct {
		msg    string
		idx    int
		hasErr bool
	}

	testCases := []testCase{
		{
			msg:    "Retrieves correct element from the head",
			idx:    0,
			hasErr: false,
		},
		{
			msg:    "Retrieves correct element from the tail",
			idx:    length - 1,
			hasErr: false,
		},
		{
			msg:    "Retrieves correct element from the middle",
			idx:    length / 2,
			hasErr: false,
		},
		{
			msg:    "Throws an error when index is negative",
			idx:    -1,
			hasErr: true,
		},
		{
			msg:    "Throws an error when index is out of bonds",
			idx:    length,
			hasErr: true,
		},
	}

	assert := assert.New(t)
	for _, test := range testCases {
		value, err := list.Get(test.idx)
		if test.hasErr {
			assert.NotNil(err, test.msg)
		} else {
			assert.Nil(err, test.msg)
			assert.Equal(getStubElement(test.idx), value, test.msg)
		}
		assert.Equal(length, list.Length(),
			"list.Get() has no side effects")
	}
}

func TestInsert(t *testing.T) {
	length := 20

	type testCase struct {
		msg    string
		idx    int
		hasErr bool
	}

	testCases := []testCase{
		{
			msg:    "Correctly inserts an element at the head",
			idx:    0,
			hasErr: false,
		},
		{
			msg:    "Correctly inserts an element at the end",
			idx:    length - 1,
			hasErr: false,
		},
		{
			msg:    "Correctly inserts an element at the middle",
			idx:    length / 2,
			hasErr: false,
		},
		{
			msg:    "Correctly inserts an element at random position",
			idx:    length / 4,
			hasErr: false,
		},
		{
			msg:    "Correctly appends an element",
			idx:    length,
			hasErr: false,
		},
		{
			msg:    "Throws an error when index is negative",
			idx:    -1,
			hasErr: true,
		},
		{
			msg:    "Throws an error when index is out of bonds",
			idx:    length + 1,
			hasErr: true,
		},
	}

	assert := assert.New(t)
	testElement := '1'
	for _, test := range testCases {
		list := getList(length)
		err := list.Insert(testElement, test.idx)

		if test.hasErr {
			assert.NotNil(err, test.msg)
			continue
		}

		//No error is produced, list length changed
		assert.Nil(err, test.msg)
		assert.Equal(length+1, list.Length(), test.msg)

		//New element is at the insert position
		value, err := list.Get(test.idx)
		assert.Nil(err, test.msg)
		assert.Equal(testElement, value, test.msg)

		//Previous element stayed in place
		if test.idx != 0 {
			previous, err := list.Get(test.idx - 1)
			assert.Nil(err, test.msg)
			assert.Equal(getStubElement(test.idx-1), previous, test.msg)
		}

		//Initial element changed position
		if test.idx != length {
			next, err := list.Get(test.idx + 1)
			assert.Nil(err, test.msg)
			assert.Equal(getStubElement(test.idx), next, test.msg)
		}
	}
}

func TestDelete(t *testing.T) {
	length := 5

	type testCase struct {
		msg    string
		idx    int
		hasErr bool
	}

	testCases := []testCase{
		{
			msg:    "Correctly deletes first element",
			idx:    0,
			hasErr: false,
		},
		{
			msg:    "Correctly deletes last element",
			idx:    length - 1,
			hasErr: false,
		},
		{
			msg:    "Correctly deletes an element in the middle",
			idx:    length / 2,
			hasErr: false,
		},
		{
			msg:    "Throws an error when index is negative",
			idx:    -1,
			hasErr: true,
		},
		{
			msg:    "Throws an error when index is out of bonds",
			idx:    length,
			hasErr: true,
		},
	}

	assert := assert.New(t)
	for _, test := range testCases {
		list := getList(length)
		err := list.Delete(test.idx)

		if test.hasErr {
			assert.NotNil(err, test.msg)
			continue
		}

		//No error produced, length changed
		assert.Nil(err, test.msg)
		assert.Equal(length-1, list.Length(), test.msg)

		//Deleted index is now occupied with next value
		if test.idx != length-1 {
			value, err := list.Get(test.idx)
			assert.Nil(err, test.msg)
			assert.Equal(getStubElement(test.idx+1), value, test.msg)
		}

		//Previous element's position is intact
		if test.idx != 0 {
			value, err := list.Get(test.idx - 1)
			assert.Nil(err, test.msg)
			assert.Equal(getStubElement(test.idx-1), value, test.msg)
		}
	}
}

func TestDeleteAll(t *testing.T) {
	type testCase struct {
		msg     string
		input   []rune
		output  []rune
		element rune
	}

	testCases := []testCase{
		{
			msg:     "Correctly deletes one middle element",
			input:   []rune{'a', 'b', 'c', 'd'},
			output:  []rune{'a', 'b', 'd'},
			element: 'c',
		},
		{
			msg:     "Correctly deletes one first element",
			input:   []rune{'a', 'b', 'c', 'd'},
			output:  []rune{'b', 'c', 'd'},
			element: 'a',
		},
		{
			msg:     "Correctly deletes one last element",
			input:   []rune{'a', 'b', 'c', 'd'},
			output:  []rune{'a', 'b', 'c'},
			element: 'd',
		},
		{
			msg:     "Correctly deletes multiple elements : 1",
			input:   []rune{'a', 'b', 'b', 'a', 'a', 'b', 'a'},
			output:  []rune{'b', 'b', 'b'},
			element: 'a',
		},
		{
			msg:     "Correctly deletes multiple elements : 2",
			input:   []rune{'b', 'a', 'b', 'a', 'b'},
			output:  []rune{'b', 'b', 'b'},
			element: 'a',
		},
		{
			msg:     "Correctly deletes all elements",
			input:   []rune{'a', 'a', 'a', 'a', 'a'},
			output:  []rune{},
			element: 'a',
		},
		{
			msg:     "Deletes no elements if those are not present",
			input:   []rune{'a', 'a', 'a', 'a', 'a'},
			output:  []rune{'a', 'a', 'a', 'a', 'a'},
			element: 'b',
		},
	}

	for _, test := range testCases {
		list := constructList(test.input)
		list.DeleteAll(test.element)
		assert.True(t, areSame(list, test.output), test.msg)
	}
}
