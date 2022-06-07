package main

import "fmt"

func main() {
	// rand.Seed(time.Now().Unix())
	// permutation := rand.Perm(1111150)
	//
	// // fmt.Println(permutation)
	// sort.InsertionSort(permutation)
	// //fmt.Println(permutation)
	//	arrSz := 10000
	arrSz := 10
	mid := arrSz / 2
	index := 0
	array := make([]int, arrSz)
	for i := 0; i < arrSz; i++ {
		if i%2 == 0 {
			array[index] = i
			index++
		} else {
			array[mid] = i
			mid++
		}
	}
	fmt.Println(array)

}
