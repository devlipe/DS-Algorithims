package sort

import "golang.org/x/exp/constraints"

// SelectionSort is an algorithm of O(n^2) complexity.
// It is in place and works by swapping the place of the min element
//
func SelectionSort[T constraints.Ordered](array []T) {
	var min_idx int

	for i := 0; i < len(array)-1; i++ {
		//we need to find the minimum element
		min_idx = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_idx] {
				min_idx = j
			}
		}
		swap(&array[min_idx], &array[i])
	}
}

func swap[T constraints.Ordered](a *T, b *T) {
	temp := *a
	*a = *b
	*b = temp
}
