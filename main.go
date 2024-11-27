package main

import (
	"fmt"

	"github.com/nearz/go-dsa/queues/ringbuffer"
)

func main() {
	r := ringbuffer.New[int](5)
	fmt.Println(r.String())
	for i := range 7 {
		err := r.Enqueue(i + 1)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println(r.String())
	v, err := r.Deque()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read: %d\n", v)
	v, err = r.Deque()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read: %d\n", v)
	v, err = r.Deque()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read: %d\n", v)
	err = r.Enqueue(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.Enqueue(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.String())
}
