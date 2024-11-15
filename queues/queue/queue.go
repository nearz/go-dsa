package queue

import (
	"errors"

	"github.com/nearz/go-dsa/lists/linkedlist"
)

type Queue[T comparable] struct {
	list *linkedlist.LinkedList[T]
}

// Queue and it's methods built on top of
// linked list package.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		list: linkedlist.New[T](),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.list.Append(val)
}

func (q *Queue[T]) Deque() (T, error) {
	v, err := q.list.RemoveAt(0)
	// If RemoveAt errors then it will be when
	// the underlying linked list is empty.
	// If so return zero value for type and
	// an error message.
	if err != nil {
		var z T
		return z, errors.New("Empty queue")
	}
	return v, nil
}

func (q *Queue[T]) Peek() (T, error) {
	v, err := q.list.Get(0)
	// If RemoveAt errors then it will be when
	// the underlying linked list is empty.
	// If so return zero value for type and
	// an error message.
	if err != nil {
		var z T
		return z, errors.New("Empty queue")
	}
	return v, nil
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

/////////////// UTILITY FUNCTIONS ///////////////

func (q *Queue[T]) String() string {
	return q.list.String()
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}
