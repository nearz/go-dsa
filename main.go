package main

import (
	"fmt"

	"github.com/nearz/go-dsa/stacks/stack"
)

func main() {
	s := stack.New[int]()
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Len())
	fmt.Println(s.Empty())

	v, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Peek: %d\n", v)

	v, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pop: %d\n", v)

	v, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pop: %d\n", v)

	v, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pop: %d\n", v)

	v, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pop: %d\n", v)

	v, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pop: %d\n", v)
}
