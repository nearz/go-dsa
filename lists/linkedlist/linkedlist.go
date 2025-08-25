package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

/*
Method list for LinkedList
Append(val T)
Add(vals ...T)
Prepend(val T)
Get(idx int) val T, err
Set(val T, idx int) err
InsertAt(val T, idx int) err
RemoveAt(idx int) val T, err
Remove(val T) val T, err
Len() int
Empty() bool
----- UTILITY FUNCTIONS -----
Values() []T
String() string
Print()
*/

// Node struct where the value is a comparable type
type Node[T comparable] struct {
	val  T
	next *Node[T]
}

// LinkedList struct to track head, tail and length
type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *LinkedList[T]) Append(val T) {
	n := &Node[T]{
		val:  val,
		next: nil,
	}
	// If list is empty, e.g head, tail = nil, point head and
	// tail to new node. Else, append to list by pointing tail.next to
	// the new node and moving tail pointer to the new node.
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.len += 1
}

func (l *LinkedList[T]) Add(vals ...T) {
	for _, v := range vals {
		l.Append(v)
	}
}

func (l *LinkedList[T]) Prepend(val T) {
	n := &Node[T]{
		val:  val,
		next: nil,
	}
	// If list is empty, e.g head, tail = nil, point head and
	// tail to new node. Else, prepend to list by setting the nodes
	// next pointer to the head and moving the head pointer to the new node.
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.len += 1
}

func (l *LinkedList[T]) Get(idx int) (T, error) {
	// Check if the list is empty or if the index
	// requested is out of bounds. If so return zero
	// value for the type and an error message.
	var z T
	if l.len == 0 {
		return z, errors.New("Empty list")
	} else if idx >= l.len || idx < 0 {
		return z, errors.New("Index out of bounds")
	}
	// Loop and increment node up until cur equals
	// the node at the index requested. Then return
	// the value of the cur node.
	cur := l.head
	for i := 0; i < idx; i, cur = i+1, cur.next {
	}
	return cur.val, nil
}

func (l *LinkedList[T]) Set(val T, idx int) error {
	// Check if the list is empty or if the index
	// requested is out of bounds. If so return error.
	if l.len == 0 {
		return errors.New("Empty list")
	} else if idx >= l.len || idx < 0 {
		return errors.New("Index out of bounds")
	}
	// Loop and increment node up until cur equals
	// the node at the index requested. Then set
	// the node's value
	cur := l.head
	for i := 0; i < idx; i, cur = i+1, cur.next {
	}
	cur.val = val
	return nil
}

func (l *LinkedList[T]) InsertAt(val T, idx int) error {
	// If length of list is 0 and the index requested
	// is 0 then add the value to the list. Otherwise
	// return the error that the list is empty and index
	// is greater than 0.
	// Also check if the request index is out of bounds.
	if l.len == 0 {
		if idx == 0 {
			l.Append(val)
			return nil
		}
		return errors.New("Empty list and index > 0")
	} else if idx > l.len || idx < 0 {
		return errors.New("Index out of bounds")
	}

	n := &Node[T]{
		val:  val,
		next: nil,
	}
	// If index requested is 0 then prepend the node to the
	// front of the list. Else if the index requested is
	// equal to the length of the list append the node to
	// the end of the list. Typically user will use those
	// functions available.
	if idx == 0 {
		l.Prepend(val)
		return nil
	} else if idx == l.len {
		l.Append(val)
		return nil
	}
	// Set cur to head and loop through until the node
	// present at the index prior to the index to be
	// inserted at. Point the current node's next to
	// the currents next and reassing current's next
	// to the new node.
	cur := l.head
	for i := 0; i < idx-1; i, cur = i+1, cur.next {
	}
	n.next = cur.next
	cur.next = n
	l.len += 1
	return nil
}

func (l *LinkedList[T]) RemoveAt(idx int) (T, error) {
	// If length of list is 0 then return zero value for type
	// and the error message.
	// Also check if the request index is out of bounds. If so
	// return the zero value of the type and the error message.
	var z T
	if l.len == 0 {
		return z, errors.New("Empty list")
	} else if idx >= l.len || idx < 0 {
		return z, errors.New("Index out of bounds")
	}
	// If index to remove is 0-head. Capture the value and
	// the head next pointer. Set the head next pointer to nil
	// then move the head pointer to the next node.
	if idx == 0 {
		val := l.head.val
		next := l.head.next
		l.head.next = nil
		l.head = next
		l.len -= 1
		if l.len == 0 {
			l.tail = nil
		}
		return val, nil
	}

	// Set a cur and prev pointer to the head. Then loop and
	// increment up to the node at the requested index. The pev
	// pointer is used to track the node prior to the one being removed
	cur := l.head
	prev := l.head
	for i := 0; i < idx; i, cur, prev = i+1, cur.next, cur {
	}
	// If the current node is the tail then set the prior node's next
	// to nil and point the tail to that previous node.
	if cur == l.tail {
		prev.next = nil
		l.tail = prev
		l.len -= 1
		return cur.val, nil
	}
	// Point the previous node's next to the current node's next
	// effectively removing the current node from the list. Then
	// if the list is only 1 node in length and removing it results
	// in an empty list, reset the head and tail to nil.
	prev.next = cur.next
	cur.next = nil
	l.len -= 1
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}
	return cur.val, nil
}

func (l *LinkedList[T]) Remove(val T) (T, error) {
	// Loop through the list until the current node's value is equal
	// to the request value or the current node's next is nil. If the
	// current node's value is equal to the requested value utilize
	// RemoveAt function to remove the node. Else the value is not found
	// in the list.
	var z T
	if l.len == 0 {
		return z, errors.New("Empty List")
	}
	i := 0
	cur := l.head
	for ; cur.val != val && cur.next != nil; i, cur = i+1, cur.next {
	}
	if cur.val == val {
		return l.RemoveAt(i)
	}
	return z, errors.New("Value not found")
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) Empty() bool {
	return l.len == 0
}

/////////////// UTILITY FUNCTIONS ///////////////

// Return the values in the list as a slice containing the values
// Was created primarly for debugging and testing
func (l *LinkedList[T]) Values() []T {
	vals := []T{}
	for n := l.head; n != nil; n = n.next {
		vals = append(vals, n.val)
	}
	return vals
}

// Generate a string representing the values in the list. E.g "1, 2, 3"
// Was created primarly for debugging and testing
func (l *LinkedList[T]) String() string {
	vals := []string{}
	for n := l.head; n != nil; n = n.next {
		vals = append(vals, fmt.Sprintf("%v", n.val))
	}
	str := strings.Join(vals, ", ")
	return str
}

// Print the values of the list.
// Was created primarly for debugging and testing
func (l *LinkedList[T]) Print() {
	if l.len == 0 {
		fmt.Println("Empty")
	}
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}
