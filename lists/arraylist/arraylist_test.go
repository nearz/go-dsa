package arraylist

import "testing"

func TestNew(t *testing.T) {
	i := New(0, 1, 2, 3)
	s := "0, 1, 2, 3"
	lstr := i.String()
	if lstr != s {
		t.Errorf("New() = %s, expected %s", lstr, s)
	}

	a := New[int]()
	if a.Len() != 0 && a.Cap() != 0 {
		t.Error("New() expected and empty arraylist")
	}
}

func TestAppend(t *testing.T) {
	// Testing appending to empty list
	i := New[int]()
	i.Append(0)
	s := "0"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Append() = %s, expected %s", lstr, s)
	}

	// Test appending val to list with existing elements
	i = New(0, 1, 2, 3, 4)
	i.Append(9)
	s = "0, 1, 2, 3, 4, 9"
	lstr = i.String()
	if lstr != s {
		t.Errorf("Append() = %s, expected %s", lstr, s)
	}
}

func TestAdd(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	s := "0, 1, 2, 3, 4"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Add() = %s, expected %s", lstr, s)
	}

	str := New[string]()
	str.Add("Go", "Python", "C")
	s = "Go, Python, C"
	lstr = str.String()
	if lstr != s {
		t.Errorf("Add() = %s, expected %s", lstr, s)
	}
}

func TestPrepend(t *testing.T) {
	// Test prepend to empty list
	i := New[int]()
	i.Prepend(0)
	s := "0"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Prepend() = %s, expected %s", lstr, s)
	}

	i = New[int]()
	i.Add(1, 2, 3)
	i.Prepend(4)
	s = "4, 1, 2, 3"
	lstr = i.String()
	if lstr != s {
		t.Errorf("Prepend() = %s, expected %s", lstr, s)
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

	// Test error. When trying to access and index that is out of bounds
	_, err = i.Get(9)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Index out of bounds")
	}

	// Test error. Trying to access and empty list
	i = New[int]()
	_, err = i.Get(0)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Empty list")
	}
}

func TestSet(t *testing.T) {
	i := New[int]()
	i.Add(0, 1, 2, 3, 4)
	err := i.Set(9, 3)
	if err != nil {
		t.Error(err)
	}
	s := "0, 1, 2, 9, 4"
	lstr := i.String()
	if lstr != s {
		t.Errorf("Add() = %s, expected %s", lstr, s)
	}

	// Test error. When trying to access and index that is out of bounds
	err = i.Set(0, 9)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Index out of bounds")
	}

	// Test error. Trying to access and empty list
	i = New[int]()
	err = i.Set(0, 0)
	if err == nil {
		t.Error("Get() err == nil, expected error message. Empty list")
	}
}

func TestInsertAt(t *testing.T) {
	// Test inserting to empty list at index 0
	i := New[int]()
	err := i.InsertAt(0, 0)
	if err != nil {
		t.Error(err)
	}
	s := "0"
	lstr := i.String()
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	// Test appending element to end of arraylist
	i.Add(1, 2, 3, 4)
	err = i.InsertAt(9, 5)
	if err != nil {
		t.Error(err)
	}
	s = "0, 1, 2, 3, 4, 9"
	lstr = i.String()
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	err = i.InsertAt(8, 2)
	if err != nil {
		t.Error(err)
	}
	s = "0, 1, 8, 2, 3, 4, 9"
	lstr = i.String()
	if lstr != s {
		t.Errorf("InsertAt() = %s, expected %s", lstr, s)
	}

	// Test index out of bounds error
	err = i.InsertAt(7, 9)
	if err == nil {
		t.Error("InserAt() err == nil, expected error message. Index out of bounds")
	}

	// Test empty list and index > 0 error
	i = New[int]()
	err = i.InsertAt(0, 1)
	if err == nil {
		t.Error("InsertAt() err == nil, expected error message. Empty list and index > 0")
	}
}

func TestRemoveAt(t *testing.T) {
	i := New(0, 1, 2, 3, 4)
	v, err := i.RemoveAt(2)
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("RemoveAt() = %d, expected 2", v)
	}

	// Test removing from last index
	v, err = i.RemoveAt(3)
	if err != nil {
		t.Error(err)
	}
	if v != 4 {
		t.Errorf("RemoveAt() = %d, expected 4", v)
	}

	// Test error index out of bounds
	_, err = i.RemoveAt(6)
	if err == nil {
		t.Error("RemoveAt() err == nil, expected error message. Index out of bounds")
	}

	// Test error empty list
	i = New[int]()
	_, err = i.RemoveAt(0)
	if err == nil {
		t.Error("RemoveAt() err == nil, expected error message. Empty list")
	}
}

func TestRemove(t *testing.T) {
	i := New(0, 1, 2, 3, 4)
	v, err := i.Remove(2)
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("Remove() = %d, expected 2", v)
	}

	// Test error value not found
	_, err = i.Remove(6)
	if err == nil {
		t.Error("Remove() err == nil, expected error message. Value not found")
	}
}

func TestResizing(t *testing.T) {
	// initial len should be 5 and
	// and cap should be double len
	i := New(0, 1, 2, 3, 4)
	l := i.Len()
	c := i.Cap()
	if l != 5 && c != 10 {
		t.Error("Array resizing not correct")
	}

	// len should increase by 1 and cap
	// should remain at 10
	i.Append(5)
	l = i.Len()
	c = i.Cap()
	if l != 6 && c != 10 {
		t.Error("Array resizing not correct")
	}

	// len should increase to 10 and cap
	// should remain at 10
	i.Add(6, 7, 8, 9)
	l = i.Len()
	c = i.Cap()
	if l != 10 && c != 10 {
		t.Error("Array resizing not correct")
	}

	// len should increase by 1 and cap
	// should double to 20
	i.Append(10)
	l = i.Len()
	c = i.Cap()
	if l != 11 && c != 20 {
		t.Error("Array resizing not correct")
	}

	// len shoud decrease by 3
	// cap remains the same
	i.RemoveAt(i.Len() - 1)
	i.RemoveAt(i.Len() - 1)
	i.RemoveAt(i.Len() - 1)
	l = i.Len()
	c = i.Cap()
	if l != 8 && c != 20 {
		t.Error("Array resizing not correct")
	}

	// len shoud decrease by 3
	// cap remains the same
	i.RemoveAt(i.Len() - 1)
	i.RemoveAt(i.Len() - 1)
	i.RemoveAt(i.Len() - 1)
	l = i.Len()
	c = i.Cap()
	if l != 5 && c != 20 {
		t.Error("Array resizing not correct")
	}

	// len shoud decrease by 1
	// cap shrinks by 0.25
	i.RemoveAt(i.Len() - 1)
	l = i.Len()
	c = i.Cap()
	if l != 4 && c != 5 {
		t.Error("Array resizing not correct")
	}
}
