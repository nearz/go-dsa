package ringbuffer

import (
	"errors"
	"fmt"
	"strings"
)

type RingBuffer[T any] struct {
	elems     []T
	write     int
	read      int
	size      int
	len       int
	full      bool
	overwrite bool
}

// Client must sepcify the size of the buffer and whether
// writing should overwrite the oldest data or stop enqueues
// when the buffer is full.
func New[T any](size int, overwrite bool) *RingBuffer[T] {
	return &RingBuffer[T]{
		elems:     make([]T, size),
		size:      size,
		full:      false,
		overwrite: overwrite,
	}
}

func (r *RingBuffer[T]) Enqueue(val T) error {
	// If client sets overwrite to true then func to overwrite
	// old data is called to enqueue the val. Else the enqueue is
	// called which will return an error if the buffer is full.
	if r.overwrite {
		r.enqueueOverwrite(val)
		return nil
	} else {
		err := r.enqueue(val)
		if err != nil {
			return err
		}
		return nil
	}
}

func (r *RingBuffer[T]) enqueue(val T) error {
	if r.full {
		return errors.New("Buffer full.")
	}
	r.elems[r.write] = val
	// Increment the write pointer and modulo wraps write to front
	// if at the end of the array
	r.write = (r.write + 1) % len(r.elems)
	// Sets full to true if write equals read
	r.full = r.write == r.read
	// if len of elements is not equal to size then incement
	// len by 1.
	if r.len != r.size {
		r.len++
	}
	return nil
}

func (r *RingBuffer[T]) enqueueOverwrite(val T) {
	r.elems[r.write] = val
	// Increment the write pointer and modulo wraps write to front
	// if at the end of the array
	r.write = (r.write + 1) % len(r.elems)
	// if len is equal to size then increment read pointer to
	// overwrite the oldest data with newest data on next write.
	// else increase the len by 1.
	if r.len == r.size {
		r.read = (r.read + 1) % len(r.elems)
	} else {
		r.len++
	}
}

func (r *RingBuffer[T]) Deque() (T, error) {
	if r.len == 0 {
		var z T
		return z, errors.New("Buffer empty.")
	}
	val := r.elems[r.read]
	// Increment the read pointer and modulo wraps read to front
	// if at the end of the array
	r.read = (r.read + 1) % len(r.elems)
	r.full = false
	r.len--
	return val, nil
}

func (r *RingBuffer[T]) Len() int {
	return r.len
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

func (r *RingBuffer[T]) Stringer() string {
	return fmt.Sprintf(
		"Overwrite: %v, Full: %v, Len: %d, Read: %d, Write: %d, Data: %v",
		r.overwrite,
		r.full,
		r.len,
		r.read,
		r.write,
		r.elems,
	)
}
