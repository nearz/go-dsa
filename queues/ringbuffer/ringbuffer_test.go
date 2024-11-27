package ringbuffer

import "testing"

func TestEnqueue(t *testing.T) {
	r := New[int](5)
	for i := range 5 {
		err := r.Enqueue(i + 1)
		if err != nil {
			t.Error(err)
			break
		}
	}
	s := "1, 2, 3, 4, 5"
	rstr := r.String()
	if rstr != s {
		t.Errorf("Enqueue() final buffer = %s, expected %s", rstr, s)
	}

	// Test error when enqueue on full ring buffer.
	err := r.Enqueue(8)
	if err == nil {
		t.Error("Enqueue() err == nil, expected error message. Buffer full.")
	}
}

func TestDeque(t *testing.T) {
	r := New[int](5)

	// Test error when deque on empty buffer
	_, err := r.Deque()
	if err == nil {
		t.Error("Deque() err == nil, expected error message. Buffer empty")
	}

	for i := range 5 {
		err := r.Enqueue(i + 1)
		if err != nil {
			t.Error(err)
			break
		}
	}

	v, err := r.Deque()
	if err != nil {
		t.Error(err)
	}
	if v != 1 {
		t.Errorf("Deque() = %d, expected 1", v)
	}

	v, err = r.Deque()
	if err != nil {
		t.Error(err)
	}
	if v != 2 {
		t.Errorf("Deque() = %d, expected 2", v)
	}

	v, err = r.Deque()
	if err != nil {
		t.Error(err)
	}
	if v != 3 {
		t.Errorf("Deque() = %d, expected 3", v)
	}
}

func TestWrap(t *testing.T) {
	r := New[int](5)
	for i := range 5 {
		err := r.Enqueue(i + 1)
		if err != nil {
			t.Error(err)
			break
		}
	}

	_, err := r.Deque()
	if err != nil {
		t.Error(err)
	}
	_, err = r.Deque()
	if err != nil {
		t.Error(err)
	}

	err = r.Enqueue(0)
	if err != nil {
		t.Error(err)
	}

	err = r.Enqueue(0)
	if err != nil {
		t.Error(err)
	}

	s := "0, 0, 3, 4, 5"
	rstr := r.String()
	if rstr != s {
		t.Errorf("Test Wrap final buffer = %s, expected %s", rstr, s)
	}
}
