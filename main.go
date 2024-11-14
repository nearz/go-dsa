package main

import (
	"fmt"

	"github.com/nearz/go-dsa/lists/arraylist"
)

func main() {
	a := arraylist.New[int]()
	err := a.InsertAt(9, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a.String())
}
