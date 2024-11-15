package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := New[int]()
	q.Enqueue(0)
	s := "0"
	qstr := q.String()
	if qstr != s {
		t.Errorf("Enqueue() queue values = %s, expected %s", qstr, s)
	}
	if q.Len() != 1 {
		t.Errorf("Queue length = %d, expected 1", q.Len())
	}

	for i := 1; i < 4; i++ {
		q.Enqueue(i)
	}

	s = "0, 1, 2, 3"
	qstr = q.String()
	if qstr != s {
		t.Errorf("Enqueue() queue values = %s, expected %s", qstr, s)
	}
	if q.Len() != 4 {
		t.Errorf("Queue length = %d, expected 4", q.Len())
	}

	sq := New[string]()
	sq.Enqueue("C")
	sq.Enqueue("Python")
	sq.Enqueue("Go")
	s = "C, Python, Go"
	qstr = sq.String()
	if qstr != s {
		t.Errorf("Enqueue() queue values = %s, expected %s", qstr, s)
	}
	if sq.Len() != 3 {
		t.Errorf("Queue length = %d, expected 3", q.Len())
	}
}

func TestDeque(t *testing.T) {
	q := New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	v, err := q.Deque()
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("Deque() = %d, expected 0", v)
	}
	if q.Len() != 2 {
		t.Errorf("Queue length = %d, expected 2", q.Len())
	}
	s := "1, 2"
	qstr := q.String()
	if qstr != s {
		t.Errorf("Deque() queue values = %s, expected %s", qstr, s)
	}

	v, err = q.Deque()
	if err != nil {
		t.Error(err)
	}
	if v != 1 {
		t.Errorf("Deque() = %d, expected 1", v)
	}
	if q.Len() != 1 {
		t.Errorf("Queue length = %d, expected 1", q.Len())
	}
	s = "2"
	qstr = q.String()
	if qstr != s {
		t.Errorf("Deque() queue values = %s, expected %s", qstr, s)
	}

	// Test error for deque from empty queue
	q = New[int]()
	_, err = q.Deque()
	if err == nil {
		t.Error("Deque() err == nil, expected error message. Empty queue")
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	v, err := q.Peek()
	if err != nil {
		t.Error(err)
	}
	if v != 0 {
		t.Errorf("Peek() = %d, expected 0", v)
	}

	// Test error for peek on empty queue
	q = New[int]()
	_, err = q.Peek()
	if err == nil {
		t.Error("Peek() err == nil, expected error message. Empty queue")
	}
}
