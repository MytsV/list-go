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

func TestClone(t *testing.T) {
	type testCase struct {
		msg    string
		length int
	}

	testCases := []testCase{
		{
			msg:    "Correctly clones a non-empty list",
			length: 15,
		},
		{
			msg:    "Correctly clones am empty list",
			length: 0,
		},
	}

	assert := assert.New(t)
	for _, test := range testCases {
		list := getList(test.length)
		clone := list.Clone()
		assert.NotSame(list, *clone, test.msg)

		assert.Equal(list.Length(), clone.Length(), test.msg)
		for i := 0; i < list.Length(); i++ {
			value, err := list.Get(i)
			valueClone, errClone := clone.Get(i)
			assert.Nil(err, test.msg)
			assert.Nil(errClone, test.msg)
			assert.Equal(value, valueClone, test.msg)
		}
	}
}

func TestReverse(t *testing.T) {
	type testCase struct {
		msg    string
		input  []rune
		output []rune
	}

	testCases := []testCase{
		{
			msg:    "Reverses a non-empty list : 1",
			input:  []rune{'a', 'b', 'c', 'd', 'e'},
			output: []rune{'e', 'd', 'c', 'b', 'a'},
		},
		{
			msg:    "Reverses a non-empty list : 2",
			input:  []rune{'a', 'a', 'b', 'b'},
			output: []rune{'b', 'b', 'a', 'a'},
		},
		{
			msg:    "Doesn't panic on empty lists",
			input:  []rune{},
			output: []rune{},
		},
	}

	for _, test := range testCases {
		list := constructList(test.input)
		list.Reverse()
		assert.True(t, areSame(list, test.output), test.msg)
	}
}

func TestFindFirst(t *testing.T) {
	type testCase struct {
		msg     string
		input   []rune
		element rune
		pos     int
	}

	testCases := []testCase{
		{
			msg:     "Finds the element in the end",
			input:   []rune{'a', 'b', 'c', 'd', 'e'},
			element: 'e',
			pos:     4,
		},
		{
			msg:     "Finds the element in the beginning",
			input:   []rune{'a', 'b', 'c', 'd', 'e'},
			element: 'a',
			pos:     0,
		},
		{
			msg:     "Finds the element in the middle",
			input:   []rune{'a', 'b', 'c', 'd', 'e'},
			element: 'c',
			pos:     2,
		},
		{
			msg:     "Finds actually the first element",
			input:   []rune{'a', 'b', 'b', 'b', 'a'},
			element: 'b',
			pos:     1,
		},
		{
			msg:     "Returns -1 if there is no such element : 1",
			input:   []rune{'a', 'b', 'b', 'b', 'a'},
			element: '9',
			pos:     -1,
		},
		{
			msg:     "Returns -1 if there is no such element : 2",
			input:   []rune{},
			element: 'a',
			pos:     -1,
		},
	}

	for _, test := range testCases {
		list := constructList(test.input)
		pos := list.FindFirst(test.element)
		assert.Equal(t, test.pos, pos, test.msg)
	}
}

func TestFindLast(t *testing.T) {
	type testCase struct {
		msg     string
		input   []rune
		element rune
		pos     int
	}

	testCases := []testCase{
		{
			msg:     "Finds actually the last element : 1",
			input:   []rune{'a', 'b', 'b', 'b', 'a'},
			element: 'b',
			pos:     3,
		},
		{
			msg:     "Finds actually the last element : 2",
			input:   []rune{'a', 'b', 'c', 'c', 'a'},
			element: 'a',
			pos:     4,
		},
		{
			msg:     "Returns -1 if there is no such element",
			input:   []rune{'a', 'b', 'c', 'd', 'e'},
			element: '1',
			pos:     -1,
		},
	}

	for _, test := range testCases {
		list := constructList(test.input)
		pos := list.FindLast(test.element)
		assert.Equal(t, test.pos, pos, test.msg)
	}
}

func TestClear(t *testing.T) {
	list := getList(15)
	list.Clear()
	assert.Equal(t, 0, list.Length(), "Makes any list empty")
}

func TestExtend(t *testing.T) {
	type testCase struct {
		msg      string
		original []rune
		extender []rune
		result   []rune
	}

	testCases := []testCase{
		{
			msg:      "Concantenates two non-empty lists",
			original: []rune{'a', 'b', 'c'},
			extender: []rune{'d', 'e'},
			result:   []rune{'a', 'b', 'c', 'd', 'e'},
		},
		{
			msg:      "Can extend with empty lists",
			original: []rune{'a', 'b', 'c'},
			extender: []rune{},
			result:   []rune{'a', 'b', 'c'},
		},
		{
			msg:      "Can extend an empty list",
			original: []rune{},
			extender: []rune{'d', 'e'},
			result:   []rune{'d', 'e'},
		},
	}

	for _, test := range testCases {
		list := constructList(test.original)
		extender := constructList(test.extender)
		list.Extend(extender)
		assert.True(t, areSame(list, test.result), test.msg)
		//No side effects on extender list
		assert.True(t, areSame(extender, test.extender), test.msg)
	}
}
