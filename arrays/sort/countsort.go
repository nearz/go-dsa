package sort

import (
	"fmt"

	au "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func CountSort(arr []int) {
	fmt.Printf("Max: %d\n", au.MaxInt(arr))
	countArr := make([]int, au.MaxInt(arr)+1)
	fmt.Printf("S: %v\n", countArr)
	for _, v := range arr {
		countArr[v]++
	}
	fmt.Println(countArr)
}
