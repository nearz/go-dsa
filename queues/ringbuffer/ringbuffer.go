package ringbuffer

import (
	"errors"
	"fmt"
	"strings"
)

type RingBuffer[T any] struct {
	elems []T
	write int
	read  int
	full  bool
}

func New[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		elems: make([]T, size),
		write: 0,
		read:  0,
		full:  false,
	}
}

func (r *RingBuffer[T]) Enqueue(val T) error {
	if r.full {
		return errors.New("Buffer full.")
	}
	r.elems[r.write] = val
	// Increment the write pointer and modulo wraps write to front
	// if at the end of the array
	r.write = (r.write + 1) % len(r.elems)
	// Sets full to true if write equals read
	r.full = r.write == r.read
	return nil
}

func (r *RingBuffer[T]) Deque() (T, error) {
	// If buffer is not full and read equals write then the buffer is empty
	if !r.full && r.read == r.write {
		var z T
		return z, errors.New("Buffer empty.")
	}
	val := r.elems[r.read]
	// Increment the read pointer and modulo wraps read to front
	// if at the end of the array
	r.read = (r.read + 1) % len(r.elems)
	r.full = false
	return val, nil
}

func (r *RingBuffer[T]) String() string {
	vals := []string{}
	rLen := len(r.elems)
	for i := 0; i < rLen; i++ {
		vals = append(vals, fmt.Sprintf("%v", r.elems[i]))
	}
	str := strings.Join(vals, ", ")
	return str
}
