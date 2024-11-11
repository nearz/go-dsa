package doublylinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// Node struct where the value is a comparable type
// includes a prev pointer to support doubly linked
// list features
type Node[T comparable] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

// DoublyLinkedList struct to track head, tail and length
type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (d *DoublyLinkedList[T]) Append(val T) {
	n := Node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}
	// If list is empty, e.g head, tail == nil, point head and
	// tail to new node. Else, append to list by first pointing
	// the new node's prev to current tail, then tail's next to
	// the new node. Finally moving tail to the new node.
	if d.head == nil && d.tail == nil {
		d.head = &n
		d.tail = &n
	} else {
		n.prev = d.tail
		d.tail.next = &n
		d.tail = &n
	}
	d.len += 1
}

func (d *DoublyLinkedList[T]) Add(vals ...T) {
	for _, v := range vals {
		d.Append(v)
	}
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
	n := Node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}
	// If list is empty, e.g head, tail == nil, point head and
	// tail to new node. Else, prepend by list by first pointing
	// the new node's next to the current head, then the current
	// head's prev to the new node. Finally moving head to the new node.
	if d.head == nil && d.tail == nil {
		d.head = &n
		d.tail = &n
	} else {
		n.next = d.head
		d.head.prev = &n
		d.head = &n
	}
	d.len += 1
}

// Get the value from the list at the index requested.
func (d *DoublyLinkedList[T]) Get(idx int) (T, error) {
	// Check if the list is empty or if the index
	// requested is out of bounds. If so return zero
	// value for the type and an error message.
	var z T
	if d.len == 0 {
		return z, errors.New("Empty list")
	} else if idx >= d.len || idx < 0 {
		return z, errors.New("Index out of bounds")
	}
	// Loop and increment node up until cur equals
	// the node at the index requested. Then return
	// the value of the cur node.
	cur := d.head
	for i := 0; i < idx; i, cur = i+1, cur.next {
	}
	return cur.val, nil
}

// Insert the a value at the requested index.
func (d *DoublyLinkedList[T]) InsertAt(val T, idx int) error {
	// If length of list is 0 and the index requested
	// is 0 then add the value to the list. Otherwise
	// return the error that the list is empty and index
	// is greater than 0.
	// Also check if the request index is out of bounds.
	if d.len == 0 {
		if idx == 0 {
			d.Append(val)
			return nil
		}
		return errors.New("Empty list and index > 0")
	} else if idx > d.len || idx < 0 {
		return errors.New("Index out of bounds")
	}

	n := Node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}
	// If index requested is 0 then prepend the node to the
	// front of the list. Else if the index requested is
	// equal to the length of the list append the node to
	// the end of the list. Typically user will use those
	// functions available.
	if idx == 0 {
		d.Prepend(val)
		return nil
	} else if idx == d.len {
		d.Append(val)
		return nil
	}
	// Set cur to head and loop through until the node
	// present at the index prior to the index to be
	// inserted at. Assign next and prev pointers
	cur := d.head
	for i := 0; i < idx-1; i, cur = i+1, cur.next {
	}
	n.next = cur.next
	n.prev = cur
	cur.next.prev = &n
	cur.next = &n
	d.len += 1
	return nil
}

// Remove the node from the list at the requested index
func (d *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	// If length of list is 0 then return zero value for type
	// and the error message.
	// Also check if the request index is out of bounds. If so
	// return the zero value of the type and the error message.
	var z T
	if d.len == 0 {
		return z, errors.New("Empty list")
	} else if idx >= d.len || idx < 0 {
		return z, errors.New("Index out of bounds")
	}

	// If index to remove is 0-head. Capture the value and
	// adjust pointers to remove the node.
	if idx == 0 {
		val := d.head.val
		next := d.head.next
		d.head.next = nil
		next.prev = nil
		d.head = next
		d.len -= 1
		return val, nil
	}

	// Loop up to the node at the requested index.
	cur := d.head
	for i := 0; i < idx; i, cur = i+1, cur.next {
	}

	// If the node is the trail capture the value
	// and adjust the pointers to remove the node.
	if cur == d.tail {
		val := cur.val
		prev := cur.prev
		cur.prev = nil
		prev.next = nil
		d.tail = prev
		d.len -= 1
		return val, nil
	}

	// Capture the value at the requested index
	// and remove the node.
	val := cur.val
	cur.next.prev = cur.prev
	cur.prev.next = cur.next
	cur.prev = nil
	cur.next = nil
	d.len -= 1
	if d.len == 0 {
		d.head = nil
		d.head = nil
	}
	return val, nil
}

// Remove a node from the list if the value matches the requested one.
func (d *DoublyLinkedList[T]) Remove(val T) (T, error) {
	// Loop through the list until the current node's value is equal
	// to the request value or the current node's next is nil. If the
	// current node's value is equal to the requested value utilize
	// RemoveAt function to remove the node. Else the value is not found
	// in the list.
	i := 0
	cur := d.head
	for ; cur.val != val && cur.next != nil; i, cur = i+1, cur.next {
	}
	if cur.val == val {
		return d.RemoveAt(i)
	}
	var z T
	return z, errors.New("Value not found")
}

func (d *DoublyLinkedList[T]) Len() int {
	return d.len
}

func (l *DoublyLinkedList[T]) Empty() bool {
	return l.len == 0
}

/////////////// UTILITY FUNCTIONS ///////////////

// Return the values in the list as a slice containing the values
// Was created primarly for debugging and testing
func (d *DoublyLinkedList[T]) Values() []T {
	vals := []T{}
	for n := d.head; n != nil; n = n.next {
		vals = append(vals, n.val)
	}
	return vals
}

// Generate a string representing the values in the list. E.g "1, 2, 3"
// Was created primarly for debugging and testing
func (d *DoublyLinkedList[T]) String() string {
	vals := []string{}
	for n := d.head; n != nil; n = n.next {
		vals = append(vals, fmt.Sprintf("%v", n.val))
	}
	str := strings.Join(vals, ", ")
	return str
}

// Generate a string representing the values in the list in reverse.
// Was created primarly for debugging and testing
func (d *DoublyLinkedList[T]) ReverseString() string {
	vals := []string{}
	for n := d.tail; n != nil; n = n.prev {
		vals = append(vals, fmt.Sprintf("%v", n.val))
	}
	str := strings.Join(vals, ", ")
	return str
}

// Print the values of the list.
// Was created primarly for debugging and testing
func (d *DoublyLinkedList[T]) Print() {
	if d.len == 0 {
		fmt.Println("Empty")
	}
	for n := d.head; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}
