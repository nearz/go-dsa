package stack

import (
	"errors"

	"github.com/nearz/go-dsa/lists/linkedlist"
)

type Stack[T comparable] struct {
	list *linkedlist.LinkedList[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		list: linkedlist.New[T](),
	}
}

func (s *Stack[T]) Push(val T) {
	s.list.Prepend(val)
}

func (s *Stack[T]) Pop() (T, error) {
	v, err := s.list.RemoveAt(0)
	if err != nil {
		var z T
		return z, errors.New("Empty stack")
	}
	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	v, err := s.list.Get(0)
	if err != nil {
		var z T
		return z, errors.New("Empty stack")
	}
	return v, nil
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Empty() bool {
	return s.list.Empty()
}

/////////////// UTILITY FUNCTIONS ///////////////

func (s *Stack[T]) String() string {
	return s.list.String()
}

func (s *Stack[T]) Values() []T {
	return s.list.Values()
}
