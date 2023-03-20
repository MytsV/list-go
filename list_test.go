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
			assert.Equal(getStubElement(test.idx), value, test.msg)
		}
		assert.Equal(length, list.Length(),
			"list.Get() has no side effects")
	}
}
