package linkedlist

import (
	"testing"
)

func TestAppend(t *testing.T) {
	l := New[int]()
	for i := range 3 {
		l.Append(i)
	}
	str := "0, 1, 2"
	lstr := l.String()
	if lstr != str {
		t.Errorf("Append() = %s, expected %s", lstr, str)
	}
	if l.Len() != 3 {
		t.Errorf("Len() = %d, expected 3", l.Len())
	}

	str = "0, 1, 2, 8"
	l.Append(8)
	lstr = l.String()
	if lstr != str {
		t.Errorf("Append() = %s, expected %s", lstr, str)
	}
}

func TestAdd(t *testing.T) {
	l := New[int]()
	l.Add(0, 1, 2)
	str := "0, 1, 2"
	lstr := l.String()
	if lstr != str {
		t.Errorf("Append() = %s, expected %s", lstr, str)
	}
}

func TestPrepend(t *testing.T) {
	l := New[int]()
	l.Add(0, 1, 2)
	str := "1, 0, 1, 2"
	l.Prepend(1)
	lstr := l.String()
	if lstr != str {
		t.Errorf("Prepend() = %s, expected %s", lstr, str)
	}
	if l.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", l.Len())
	}

	l = New[int]()
	l.Prepend(8)
	str = "8"
	lstr = l.String()
	if lstr != str {
		t.Errorf("Prepend() = %s, expected %s", lstr, str)
	}
}

func TestGet(t *testing.T) {
	// Test typical operation
	s := New[string]()
	s.Add("Golang", "Python", "C", "Zig")
	g, err := s.Get(2)
	if err != nil {
		t.Error(err)
	}
	if g != "C" {
		t.Errorf("Get()= %s, expected C", g)
	}

	// Test get on list of 1 node
	i := New[int]()
	i.Append(1)
	r, err := i.Get(0)
	if err != nil {
		t.Error(err)
	}
	if r != 1 {
		t.Errorf("Get()= %d, expected 1", r)
	}

	// Test get last node
	i = New[int]()
	i.Add(0, 1, 2, 3, 4, 5)
	r, err = i.Get(5)
	if err != nil {
		t.Error(err)
	}
	if r != 5 {
		t.Errorf("Get() = %d, expected 5", r)
	}

	// Test Error - Empty List
	b := New[bool]()
	_, err = b.Get(0)
	if err == nil {
		t.Error("Get() = nil, expected an error message. Empty List")
	}

	// Test Error - Index out of bounds
	f := New[float64]()
	f.Add(1.2, 2.1)
	_, err = f.Get(2)
	if err == nil {
		t.Error("Get() err == nil, expected an error message. Index out of bounds")
	}
}

func TestSet(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	s := "0, 1, 9, 3, 4"
	err := i.Set(9, 2)
	lstr := i.String()
	if err != nil {
		t.Error(err)
	}
	if lstr != s {
		t.Errorf("Set() = %s, expected %s", lstr, s)
	}

	err = i.Set(9, 7)
	if err == nil {
		t.Error("Set() err == nil, expected an error message. Index out of bounds")
	}

	i = New[int]()
	err = i.Set(0, 0)
	if err == nil {
		t.Error("Set() err == nil, expected an error message. Empty list")
	}
}

func TestInsertAt(t *testing.T) {
	// Test typical operation
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	s := "0, 1, 2, 8, 3, 4"
	err := i.InsertAt(8, 3)
	lstr := i.String()
	if err != nil {
		t.Error(err)
	}
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}
	if i.Len() != 6 {
		t.Errorf("Len() = %d, expected 6", i.Len())
	}

	// Test inserting at index 0, aka Prepend
	i = New[int]()
	i.Add(0, 1)
	s = "1, 0, 1"
	err = i.InsertAt(1, 0)
	lstr = i.String()
	if err != nil {
		t.Error(err)
	}
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	// Test inserting at end creating new index, aka Append
	i = New[int]()
	i.Add(0, 1)
	s = "0, 1, 1"
	err = i.InsertAt(1, 2)
	lstr = i.String()
	if err != nil {
		t.Error(err)
	}
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	// Test inserting to empty list, Index must be 0
	i = New[int]()
	err = i.InsertAt(1, 0)
	s = "1"
	lstr = i.String()
	if err != nil {
		t.Error(err)
	}
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	// Test Error - inserting to empty list and index is >0
	i = New[int]()
	err = i.InsertAt(1, 1)
	if err == nil {
		t.Error("InsertAt() = nil, expected error message. Empty list and index > 0")
	}

	// Test Error - Index out of bounds
	i = New[int]()
	i.Add(0, 1, 2, 3)
	err = i.InsertAt(0, 8)
	if err == nil {
		t.Error("InsertAt() = nil, expected error message. Index out of bounds")
	}
}

func TestRemoveAt(t *testing.T) {
	// Test typical operation
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	v, err := i.RemoveAt(2)
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("RemoveAt() = %d, expected 2", v)
	}
	s := "0, 1, 3, 4"
	lstr := i.String()
	if lstr != s {
		msg := "RemoveAt(), linked list after remove at index is not correct"
		t.Errorf("%s, result = %s, expected %s", msg, lstr, s)
	}
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}

	// Test if index to remove is 0
	v, err = i.RemoveAt(0)
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("RemoveAt() = %d, expected 0", v)
	}
	s = "1, 3, 4"
	lstr = i.String()
	if lstr != s {
		msg := "RemoveAt(), linked list after remove at index is not correct"
		t.Errorf("%s, result = %s, expected %s", msg, lstr, s)
	}
	if i.Len() != 3 {
		t.Errorf("Len() = %d, expected 3", i.Len())
	}

	// Test Error - removing with empty list
	i = New[int]()
	_, err = i.RemoveAt(1)
	if err == nil {
		t.Error("RemoveAt() - err == nil, expected error message. Empty list and index > 0")
	}

	// Test Error - Index out of bounds
	i = New[int]()
	i.Add(0, 1, 2, 3)
	_, err = i.RemoveAt(8)
	if err == nil {
		t.Error("RemoveAt() - err == nil, expected error message. Index out of bounds")
	}
}

func TestRemove(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	v, err := i.Remove(3)
	if err != nil {
		t.Error(err)
	}
	if v != 3 {
		t.Errorf("Remove() = %d, expected 3", v)
	}
	s := "0, 1, 2, 4"
	lstr := i.String()
	if lstr != s {
		msg := "Remove(), linked list after remove is not correct"
		t.Errorf("%s, result = %s, expected %s", msg, lstr, s)
	}
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}

	_, err = i.Remove(8)
	if err == nil {
		t.Error("Remove() - err == nil, expected error message. Value not found")
	}
}
