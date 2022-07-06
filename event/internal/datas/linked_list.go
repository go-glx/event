package datas

type (
	LinkedList[T comparable] struct {
		head *node[T]
		tail *node[T]
	}

	node[T comparable] struct {
		next  *node[T]
		value T
	}
)

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func newNode[T comparable](value T) *node[T] {
	return &node[T]{
		next:  nil,
		value: value,
	}
}

// Clear complexity: O(1)
func (lst *LinkedList[T]) Clear() {
	lst.head = nil
	lst.tail = nil
}

// Prepend complexity: O(1)
func (lst *LinkedList[T]) Prepend(value T) {
	node := newNode(value)
	node.next = lst.head

	lst.head = node
}

// Append complexity: O(1)
func (lst *LinkedList[T]) Append(value T) {
	// case: empty list
	if lst.tail == nil {
		node := newNode(value)
		lst.head = node
		lst.tail = node
		return
	}

	// case: append to end
	node := newNode(value)
	lst.tail.next = node
	lst.tail = node

	_ = lst
}

// Iterate complexity: O(n)
func (lst *LinkedList[T]) Iterate(fn func(T)) {
	node := lst.head
	for node != nil {
		fn(node.value)
		node = node.next
	}
}

// Len complexity: O(n)
func (lst *LinkedList[T]) Len() int {
	count := 0
	lst.Iterate(func(_ T) {
		count++
	})
	return count
}

// Remove complexity: O(n)
// This will remove first occurrence of value (if list have many copies of this value)
func (lst *LinkedList[T]) Remove(value T) {
	var prev *node[T]
	curr := lst.head

	for curr != nil {
		if curr.value != value {
			prev = curr
			curr = curr.next
			continue
		}

		switch {
		case prev == nil && curr.next == nil:
			lst.removeSingle()
		case prev == nil:
			lst.removeHead()
		case curr.next == nil:
			lst.removeTail(prev)
		default:
			lst.removeNodeBetween(prev, curr.next)
		}
		return
	}
}

// guaranteed only one element in list
// head and tail reference it
func (lst *LinkedList[T]) removeSingle() {
	lst.head = nil
	lst.tail = nil
}

// guaranteed at least 2 elements in list
func (lst *LinkedList[T]) removeHead() {
	lst.head = lst.head.next
}

// guaranteed at least 2 elements in list
func (lst *LinkedList[T]) removeTail(prev *node[T]) {
	prev.next = nil
	lst.tail = prev
}

// guaranteed at least 3 elements in list
func (lst *LinkedList[T]) removeNodeBetween(prev, next *node[T]) {
	prev.next = next
}
