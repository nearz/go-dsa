package main

import (
	"fmt"

	"github.com/nearz/go-dsa/queues/ringbuffer"
)

func main() {
	r := ringbuffer.New[int](5, true)
	fmt.Println(r.String())
	for i := range 7 {
		err := r.Enqueue(i + 1)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println(r.Stringer())
	v, err := r.Deque()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read: %d\n", v)
}
