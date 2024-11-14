package arraylist

import (
	"errors"
	"fmt"
	"strings"
)

const (
	expFactor    float32 = 2.0
	shrinkFactor float32 = 0.25
)

/*
Method list for ArrayList
Append(val T)
Add(vals ...T)
Prepend(val T)
Get(idx int) val T, err
Set(val T, idx int) err
InsertAt(val T, idx int) err
RemoveAt(idx int) val T, err
Remove(val T) val T, err
Empty()
Len()
Cap()
----- PRIVATE FUNCTIONS -----
shiftLeft(idx int)
resize(len, cap int)
shrink()
expand(n int)
----- UTILITY FUNCTIONS -----
Values() []T
String() string
*/

type ArrayList[T comparable] struct {
	elems []T
}

func New[T comparable](vals ...T) *ArrayList[T] {
	a := &ArrayList[T]{}
	if len(vals) > 0 {
		a.Add(vals...)
	}
	return a
}

func (a *ArrayList[T]) Append(val T) {
	// Increase the len to add the
	// new index at the end of the list
	// cap will increase if needed.
	a.expand(1)
	idx := len(a.elems) - 1
	a.elems[idx] = val
}

// Add 1 or more values to an arraylist
func (a *ArrayList[T]) Add(vals ...T) {
	l := len(a.elems)
	n := len(vals)
	// Adjust len of slice for new values
	// and adjust capacity if needed. Then
	// add new values to the array.
	a.expand(n)
	for i, v := range vals {
		a.elems[l+i] = v
	}
}

func (a *ArrayList[T]) Prepend(val T) {
	// Increase len by one and shift
	// all element right from index 0.
	// cap will increase if needed.
	a.expand(1)
	a.shiftRight(0)
	a.elems[0] = val
}

func (a *ArrayList[T]) Get(idx int) (T, error) {
	// Check if array is empty or if requested
	// index is out of bounds. If true return
	// the zero value of the type and error message.
	var z T
	aLen := len(a.elems)
	if aLen == 0 {
		return z, errors.New("Empty list")
	} else if idx >= aLen || idx < 0 {
		return z, errors.New("Index out of bounds")
	}
	return a.elems[idx], nil
}

func (a *ArrayList[T]) Set(val T, idx int) error {
	// Check if array is empty or if requested
	// index is out of bounds. If true return
	// error message
	aLen := len(a.elems)
	if aLen == 0 {
		return errors.New("Empty list")
	} else if idx >= aLen || idx < 0 {
		return errors.New("Index out of bounds")
	}
	a.elems[idx] = val
	return nil
}

func (a *ArrayList[T]) InsertAt(val T, idx int) error {
	aLen := len(a.elems)
	// If the list is empty and index requested is 0
	// then append the value to the list at index 0.
	// If list is empty and index is > 0 then return
	// error message. Also check if index is out of
	// bounds and return error if so.
	if aLen == 0 {
		if idx == 0 {
			a.Append(val)
			return nil
		}
		return errors.New("Empty list and index > 0")
	} else if idx > aLen || idx < 0 {
		return errors.New("Index out of bounds")
	}
	// If index requested is equal to arraylist len
	// then append the value to the list
	// no shift right required
	if idx == aLen {
		a.Append(val)
		return nil
	}
	// expand the arraylist by 1 and shift elems
	// to the right from the index and insert the
	// value at the requested index.
	a.expand(1)
	a.shiftRight(idx)
	a.elems[idx] = val
	return nil
}

func (a *ArrayList[T]) RemoveAt(idx int) (T, error) {
	// If length of list is 0 then return zero value for type
	// and the error message.
	// Also check if the request index is out of bounds. If so
	// return the zero value of the type and the error message.
	var z T
	aLen := len(a.elems)
	if aLen == 0 {
		return z, errors.New("Empty list")
	} else if idx >= aLen || idx < 0 {
		return z, errors.New("Index out of bounds")
	}
	// If index requested is the last index then capture the
	// value and remove the last index with no shift left required.
	// Shrink will adjust capacity if needed.
	if idx == aLen-1 {
		val := a.elems[aLen-1]
		a.elems = a.elems[:aLen-1]
		a.shrink()
		return val, nil
	}
	// Capture the value at the index and shift left removing
	// the index requested. Shrink capacity if needed.
	val := a.elems[idx]
	a.shiftLeft(idx)
	a.shrink()
	return val, nil
}

func (a *ArrayList[T]) Remove(val T) (T, error) {
	// Check if the arraulist is empty. if so
	// return the zero value of the type and
	// and error message
	var z T
	aLen := len(a.elems)
	if aLen == 0 {
		return z, errors.New("Empty list")
	}
	// loop over the arraylist, if value is found
	// mark found true and break.
	i := 0
	found := false
	for ; i < aLen; i++ {
		if a.elems[i] == val {
			found = true
			break
		}
	}
	if found {
		return a.RemoveAt(i)
	}
	// If element at i is not the value return
	// error that the requested value was not found.
	return z, errors.New("Value not found")
}

func (a *ArrayList[T]) Empty() bool {
	return len(a.elems) == 0
}

func (a *ArrayList[T]) Len() int {
	return len(a.elems)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.elems)
}

// ///////// PRIVATE FUNCTIONS ///////////////

// will shift all elements left from the requested
// index to the end of the list. Used when an element
// is removed from the arraylist.
func (a *ArrayList[T]) shiftLeft(idx int) {
	i := idx + 1
	aLen := len(a.elems)
	for ; i < aLen; i++ {
		a.elems[i-1] = a.elems[i]
	}
	a.elems = a.elems[:aLen-1]
}

// will shift all elements right from the requested
// index to the end of the list. This will create
// a space for a new item to be inserted. Arraylist
// must be resized first to include a new index at
// the end to shift right into.
func (a *ArrayList[T]) shiftRight(idx int) {
	i := len(a.elems) - 1
	for ; i > idx; i-- {
		a.elems[i] = a.elems[i-1]
	}
}

// will resize the arraulist after expanding
// or shrinking the array.
func (a *ArrayList[T]) resize(len, cap int) {
	newElems := make([]T, len, cap)
	for i, v := range a.elems {
		newElems[i] = v
	}
	a.elems = newElems
}

// Check if the capacity should shrink after
// removing elements. If len is less than the current
// capacity by the shrink factor then resize.
func (a *ArrayList[T]) shrink() {
	newCap := int(float32(cap(a.elems)) * shrinkFactor)
	aLen := len(a.elems)
	if aLen <= newCap {
		a.resize(aLen, newCap)
	}
}

// Will increase the len for new elements to be added
// to the arraylist. If the new len is greater than
// the current capacity, then capacity will be resized
// by the expand factor.
func (a *ArrayList[T]) expand(n int) {
	curCap := cap(a.elems)
	newLen := len(a.elems) + n
	if newLen > curCap {
		newCap := int(float32(curCap+n) * expFactor)
		a.resize(newLen, newCap)
	} else {
		a.elems = a.elems[:newLen]
	}
}

// ///////////// UTILITY FUNCTIONS ///////////////
func (a *ArrayList[T]) Values() []T {
	return a.elems
}

func (a *ArrayList[T]) String() string {
	vals := []string{}
	aLen := len(a.elems)
	for i := 0; i < aLen; i++ {
		vals = append(vals, fmt.Sprintf("%v", a.elems[i]))
	}
	str := strings.Join(vals, ", ")
	return str
}
