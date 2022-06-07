package sort

import (
	"golang.org/x/exp/constraints"
)

// merge sort is a not in place algorithm
// it has the complexity of O(nlogn),
// it is also a stable sort algorithm

func merge[T constraints.Ordered](array []T, begin, mid, end int) {
	sz := end - begin // the size of the sub array
	tempArr := make([]T, sz)
	indexFirstArr := begin
	indexSecondArr := mid
	indexTemp := 0

	for indexFirstArr < mid && indexSecondArr < end { // Here we loop through the arrays and
		// the loop ends on the size of the minor array
		// We copy the minor number to the temp array
		if array[indexFirstArr] <= array[indexSecondArr] {
			tempArr[indexTemp] = array[indexFirstArr]
			indexFirstArr++
			indexTemp++
		} else {
			tempArr[indexTemp] = array[indexSecondArr]
			indexTemp++
			indexSecondArr++
		}

	} // Now we have to copy the other elements of the bigger array
	for indexFirstArr < mid {
		tempArr[indexTemp] = array[indexFirstArr]
		indexFirstArr++
		indexTemp++
	}
	for indexSecondArr < end {
		tempArr[indexTemp] = array[indexSecondArr]
		indexTemp++
		indexSecondArr++
	}
	// now we have to copy the elements of the temp to the original array
	for k := 0; k < sz; k++ {
		array[begin+k] = tempArr[k]
	}
}

// we need to create the recursive function to split the array into sub arrays
func mergeSortImp[T constraints.Ordered](array []T, begin int, end int) {
	// test base case so the function doesn't create a loop
	if begin >= end-1 {
		return
	}
	// divide the array into two sub arrays
	mid := (begin + end) / 2
	mergeSortImp(array, begin, mid) // sort array from begin to mid
	mergeSortImp(array, mid, end)   // sort array from mid to end
	merge(array, begin, mid, end)   // merge the entire array into one

}

// This function is a gateway so that the user doesn't have to pass the array
// length
func MergeSort[T constraints.Ordered](array []T) {
	mergeSortImp(array, 0, len(array))
}
