package sort

import (
	au "github.com/nearz/go-dsa/arrays/arrayUtils"
)

func CountSort(arr []int) []int {
	countArr := make([]int, au.MaxInt(arr)+1)
	for _, v := range arr {
		countArr[v]++
	}
	cumulative := 0
	for i, v := range countArr {
		cumulative = v + cumulative
		countArr[i] = cumulative
	}
	final := make([]int, len(arr))
	for _, v := range arr {
		j := countArr[v] - 1
		final[j] = v
		countArr[v] = j
	}
	return final
}
