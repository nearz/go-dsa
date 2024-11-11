package arrays

import (
	"math"
)

/*
Two Crystal Ball Problem
Given two crystal balls that will break if dropped from a high enough distance, determine the
exact spot in which it will break in the most optimized way.

example input: [false, false, false, true, true]
the first "true" is the point at which a crystall ball will break.

CrystalBall() was my intial solution but it felt a bit messy, but is still O(sqrt(N)). Listened
to someone explain their solution which was much cleaner so decided to implement it here in go in
CrystalBallTwo().
*/

func CrystalBall(arr []bool) int {
	l := len(arr)
	// j is the jump amount, which is the sqrt of the length of the array
	j := int(math.Floor(math.Sqrt(float64(l))))
	i := 0
	f := false
	// loop through the arr by the jump amount
	for k := j; k < l; k += j {
		// use your first crystal ball to test if k is a breaking point
		if arr[k] {
			// if a breaking point is found flag it as found then
			// go back to the last jumped-to index+1 or the beginning and scan for
			// a possible prior breaking point with the second crystal ball
			// return the index of the exact breaking point
			f = true
			for ; i <= k; i++ {
				if arr[i] {
					return i
				}
			}
		}
		// record each jump-to index to go back to if needed
		i = k + 1
	}
	// if a breaking point is not found when jumping through the array
	// then scan through the array from the last jumped-to index and search
	// for a possible breaking point and return the index if found.
	if !f {
		for ; i < l; i++ {
			if arr[i] {
				return i
			}
		}
	}
	// return -1 if no breaking point index is found
	return -1
}

func CrystalBallTwo(arr []bool) int {
	l := len(arr)
	// j is the jump amount which is the sqrt of the length of the array
	j := int(math.Floor(math.Sqrt(float64(l))))
	i := j
	// loop through the array by the jump amount
	for ; i < l; i += j {
		// use your first crystall ball to test if i is a breaking point
		// if it is break out of the loop
		if arr[i] {
			break
		}
	}
	// set k to the last jumped-to point or the beginning of the array
	k := i - j
	// scan up to the last checked point of end of array to find a possible
	// prior breaking point and return the index if found
	for ; k <= i && k < l; k++ {
		if arr[k] {
			return k
		}
	}
	// return -1 if no breaking point index is found
	return -1
}

/*
A search that uses a similar method as the crystal ball problem if the array is sorted
Note: not better than Binary search for a sorted array
*/
func BlockSearch(arr []int, v int) int {
	l := len(arr)
	j := int(math.Floor(math.Sqrt(float64(l))))
	k := j
	for ; k < l; k += j {
		if arr[k] == v {
			return k
		} else if arr[k] > v {
			break
		}
	}

	i := k - j
	for ; i < k && i < l; i++ {
		if arr[i] == v {
			return i
		}
	}

	return -1
}
