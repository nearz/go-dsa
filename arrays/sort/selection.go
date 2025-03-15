package sort

func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		// set min index to i
		min := i
		// loop through unsorted partition to search for the index
		// containing the minimum value
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// if an index was found with a value less than what is
		// at i then swap it. Else arr[i] is already the min.
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
