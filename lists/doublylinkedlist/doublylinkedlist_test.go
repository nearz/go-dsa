package doublylinkedlist

import "testing"

func TestAppend(t *testing.T) {
	i := New[int]()
	i.Append(0)
	i.Append(1)
	i.Append(2)
	i.Append(3)
	s := "0, 1, 2, 3"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Append() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "3, 2, 1, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("Append() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}
}

func TestAdd(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3)
	s := "0, 1, 2, 3"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Add() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "3, 2, 1, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("Add() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}
}

func TestPrepend(t *testing.T) {
	i := New[int]()
	i.Prepend(0)
	i.Prepend(1)
	i.Prepend(2)
	i.Prepend(3)
	s := "3, 2, 1, 0"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Append() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "0, 1, 2, 3"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("Append() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}
}

func TestGet(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	v, err := i.Get(2)
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("Get() = %d, expected 2", v)
	}

	_, err = i.Get(8)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Index out of bounds")
	}

	i = New[int]()
	_, err = i.Get(0)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Empty list")
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
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	err := i.InsertAt(8, 2)
	if err != nil {
		t.Error(err)
	}
	s := "0, 1, 8, 2, 3, 4"
	lstr := i.String()
	if lstr != s {
		t.Errorf("InsertAt() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "4, 3, 2, 8, 1, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("InsertAt() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 6 {
		t.Errorf("Len() = %d, expected 6", i.Len())
	}

	// Test insert at index 0
	err = i.InsertAt(0, 0)
	if err != nil {
		t.Error(err)
	}
	s = "0, 0, 1, 8, 2, 3, 4"
	lstr = i.String()
	if lstr != s {
		t.Errorf("InsertAt() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "4, 3, 2, 8, 1, 0, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("InsertAt() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 7 {
		t.Errorf("Len() = %d, expected 7", i.Len())
	}

	// Test insert at end in new index
	err = i.InsertAt(0, i.Len())
	if err != nil {
		t.Error(err)
	}
	s = "0, 0, 1, 8, 2, 3, 4, 0"
	lstr = i.String()
	if lstr != s {
		t.Errorf("InsertAt() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "0, 4, 3, 2, 8, 1, 0, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("InsertAt() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 8 {
		t.Errorf("Len() = %d, expected 8", i.Len())
	}

	// Test error index out of bounds
	err = i.InsertAt(0, 11)
	if err == nil {
		t.Error("InsertAt() err == nil, expected error message. Index out of bounds")
	}

	// Test error empty list and index > 0
	i = New[int]()
	err = i.InsertAt(1, 1)
	if err == nil {
		t.Error("InsertAt() err == nil, expected error message. Empty list and index > 0")
	}

	// Test empty list and insert at index 0
	err = i.InsertAt(1, 0)
	if err != nil {
		t.Error(err)
	}
	s = "1"
	lstr = i.String()
	if lstr != s {
		t.Errorf("InsertAt() final list = %s, expected %s", lstr, s)
	}
}

func TestRemoveAt(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	v, err := i.RemoveAt(2)
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("RemoveAt() = %d, expected 2", v)
	}
	// Test resulting list is correct
	s := "0, 1, 3, 4"
	lstr := i.String()
	if lstr != s {
		t.Errorf("RemoveAt() final list = %s, expected %s", lstr, s)
	}
	// Test node.prev is being set right
	s = "4, 3, 1, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("RemoveAt() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}

	// Test remove at head
	v, err = i.RemoveAt(0)
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("RemoveAt() = %d, expected 0", v)
	}

	// Test remove at tail
	v, err = i.RemoveAt(i.Len() - 1)
	if err != nil {
		t.Error(err)
	}
	if v != 4 {
		t.Errorf("RemoveAt() = %d, expected 4", v)
	}

	// Test index out of bounds
	_, err = i.RemoveAt(9)
	if err == nil {
		t.Error("RemoveAt() err == nil, expected error message. Index out of bounds")
	}

	// Test remove from empty list
	i = New[int]()
	_, err = i.RemoveAt(0)
	if err == nil {
		t.Error("RemoveAt() err == nil, expected error message. Empty list")
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
	// Test resulting list is correct
	s := "0, 1, 2, 4"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Remove() final list = %s, expected %s", lstr, s)
	}
	// Test that prev pointers are set porperly
	s = "4, 2, 1, 0"
	lstr = i.ReverseString()
	if lstr != s {
		t.Errorf("Remove() test node prev pointer = %s, expected %s", lstr, s)
	}
	// Test length is being tracked properly
	if i.Len() != 4 {
		t.Errorf("Len() = %d, expected 4", i.Len())
	}

	// Test value not found error
	_, err = i.Remove(8)
	if err == nil {
		t.Error("Remove() err == nil, expected error message. Value not found")
	}
}
