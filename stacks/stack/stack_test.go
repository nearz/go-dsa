package stack

import "testing"

func TestPush(t *testing.T) {
	i := New[int]()
	i.Push(0)
	if i.Len() != 1 {
		t.Errorf("Push() length = %d, expected 1", i.Len())
	}

	i.Push(1)
	i.Push(2)
	i.Push(3)
	s := "3, 2, 1, 0"
	sstr := i.String()
	if sstr != s {
		t.Errorf("Push() stack values = %s, expected %s", sstr, s)
	}
	if i.Len() != 4 {
		t.Errorf("Push() length = %d, expected 4", i.Len())
	}
}

func TestPop(t *testing.T) {
	i := New[int]()
	i.Push(0)
	i.Push(1)
	i.Push(2)

	v, err := i.Pop()
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("Pop() = %d, expected 2", v)
	}

	v, err = i.Pop()
	if err != nil {
		t.Error(err)
	}
	if v != 1 {
		t.Errorf("Pop() = %d, expected 1", v)
	}

	v, err = i.Pop()
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("Pop() = %d, expected 0", v)
	}

	_, err = i.Pop()
	if err == nil {
		t.Error("Pop() err == nil, expected error message. Empty stack")
	}
}

func TestPeek(t *testing.T) {
	i := New[int]()
	v, err := i.Peek()
	if err == nil {
		t.Error("Peek() err == nil, expected error message. Empty stack")
	}

	i.Push(0)
	v, err = i.Peek()
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("Peek() = %d, expected 0", v)
	}
}
