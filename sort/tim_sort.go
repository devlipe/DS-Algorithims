// Tim sort algorithm
// Felipe P. Ferreira
// 09/06/2022
package sort

import (
	"golang.org/x/exp/constraints"
)

/*
	Tim sort is used languages like Java and Python.
	It uses elements form merge sort and insertion sort
	-  In-situ algorithm
	- O(nlogn)
	- Stable algorithm
*/

const RUN = 32 // first we define the size of the sub arrays

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TimSort[T constraints.Ordered](array []T) {
	sz := len(array)
	for i := 0; i < sz; i += RUN {
		InsertionSort(array[i:min((i+RUN-1), (sz))])
	}

	for size := RUN; size < sz; size *= 2 {
		for left := 0; left < sz; left += 2 * size {
			mid := left + size - 1
			right := min((left + 2*size - 1), (sz - 1))
			if mid < right {
				Merge(array, left, mid, right)
			}
		}
	}
}
