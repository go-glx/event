package datas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	lst := NewLinkedList[int]()
	testAssertLst(t, lst, []int{})

	// append
	lst.Append(4)
	lst.Append(8)
	lst.Append(15) // duplicate: 1
	lst.Append(23)
	lst.Append(15) // duplicate: 2
	lst.Append(42)
	testAssertLst(t, lst, []int{4, 8, 15, 23, 15, 42})
	assert.Equal(t, 6, lst.Len())

	// remove first occurrence
	lst.Remove(15)
	testAssertLst(t, lst, []int{4, 8, 23, 15, 42})
	assert.Equal(t, 5, lst.Len())

	// prepend
	lst.Prepend(99)
	lst.Prepend(15)
	lst.Prepend(99)
	testAssertLst(t, lst, []int{99, 15, 99, 4, 8, 23, 15, 42})
	assert.Equal(t, 8, lst.Len())

	// remove some numbers (exist)
	lst.Remove(99)
	lst.Remove(8)
	lst.Remove(42)
	testAssertLst(t, lst, []int{15, 99, 4, 23, 15})

	// remove unknown numbers
	lst.Remove(10)
	lst.Remove(5)
	lst.Remove(5)
	testAssertLst(t, lst, []int{15, 99, 4, 23, 15})

	// clear
	lst.Clear()
	testAssertLst(t, lst, []int{})
	assert.Equal(t, 0, lst.Len())
}

func testAssertLst[T comparable](t *testing.T, lst *LinkedList[T], expected []T) {
	assert.Equal(t, expected, testDumpLinkedList(lst))
}

func testDumpLinkedList[T comparable](lst *LinkedList[T]) []T {
	data := make([]T, 0)

	lst.Iterate(func(element T) {
		data = append(data, element)
	})

	return data
}
