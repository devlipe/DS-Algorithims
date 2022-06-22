package sort

import (
	"golang.org/x/exp/constraints"
)

// BinarySearch searches for the specified object in the sorted array.
// The array must be sorted in ascending order.
// If it is not sorted, the func will fail.
// If the object is not found in the array, the func will return -1.
func BinarySearch[T constraints.Ordered](array []T, key T) int {
	return binarySearchImpl(array, key, 0, len(array)-1)
}

func binarySearchImpl[T constraints.Ordered](array []T, key T, low, high int) int {
	if low > high {
		return -1 // not found
	}

	mid := (high-low)/2 + low
	if array[mid] == key {
		return mid
	}
	if array[mid] < key {
		return binarySearchImpl(array, key, mid+1, high)
	}
	return binarySearchImpl(array, key, low, mid-1)
}

func BinarySearchInterative[T constraints.Ordered](array []T, key T) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (high-low)/2 + low
		if array[mid] == key {
			return mid
		}
		if array[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1

}

func SequentialSearch[T constraints.Ordered](array []T, key T) int {
	for i, v := range array {
		if v == key {
			return i
		}
	}
	return -1
}
