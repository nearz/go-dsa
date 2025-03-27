package sort

func InsertionSort(arr []int) {
	l := len(arr)
	// loop up to 2nd to last index
	for i := 0; i < l-1; i++ {
		// set comparison key to index + 1
		k := arr[i+1]
		j := i
		// loop down to 0 or if k >= value at j
		// shift any value greater than k up 1 index
		for j > -1 && k < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		// place k in its final sorted position
		arr[j+1] = k
	}
}
