package sort

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

/*this function will take an element from the array and make it the pivot
than it will place elements that are less than on the left and the ones that are greater
on the right*/

func Partition[T constraints.Ordered](array []T, begin, end int) int {
	// the pivot must be inside the array range
	// so if begin is 10 and end is 15 we neet a number between [10 and 14]
	// so we generate a random number up to the diferente of end and begin and sum it to begin
	indexPivot := (rand.Intn((end - begin)) + begin) // Here we get a random element to be the pivot
	pivotValue := array[indexPivot]
	swap(&array[indexPivot], &array[end-1]) // Put the pivo at the end of the array

	pos := begin // when we find an element that is less than pivot we will put it in position pos

	for i := begin; i < end-1; i++ {
		if array[i] < pivotValue {
			swap(&array[pos], &array[i])
			pos++
		}
	}
	swap(&array[pos], &array[end-1]) // put pivot after the greatest element that is less than pivot
	return pos                       // return the position of the pivot

}

func QuickSortImp[T constraints.Ordered](array []T, begin, end int) {
	rand.Seed(time.Now().Unix())
	if begin == end {
		return // If begin is equal to end, than the vector has 1 element and is Ordered
	}
	pos := Partition(array, begin, end)
	QuickSortImp(array, begin, pos)
	QuickSortImp(array, pos+1, end)

}

/*
	The QuickSort method have this caracteristics
	Divider and conquer algorithm
	It is an in-place algorithm
	It is not stable, so elements that are considered equal can change orders
	@Complexity
	It's complexity is O(nlogn) -- overall. Worst complexity is O(n^2)
	This chan happen if we choose the greatest or smallest element of the sub array
	The best complexity is O(n), if we pick the middle element of the sub array
*/
func QuickSort[T constraints.Ordered](array []T) {
	QuickSortImp(array, 0, len(array))
}
