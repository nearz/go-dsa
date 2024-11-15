package main

import (
	"fmt"

	"github.com/nearz/go-dsa/queues/queue"
)

func main() {
	q := queue.New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	v, err := q.Peek()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	v, err = q.Deque()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	v, err = q.Deque()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	v, err = q.Deque()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
