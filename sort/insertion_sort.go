package sort

import "golang.org/x/exp/constraints"

/*
*	Insertion sort is an algorithm with a complexity of O(n^2)
 */
func InsertionSort[T constraints.Ordered](array []T) {

	for i := 1; i < len(array); i++ {
		// We select the i element of the array
		// Than we compare it with the i elements that are first of it
		key := array[i]

		//The elements that are greater than the key, we move up
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}
