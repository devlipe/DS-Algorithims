package main

import "fmt"

func main() {
	// rand.Seed(time.Now().Unix())
	// permutation := rand.Perm(50)
	//
	// fmt.Println(permutation)
	// sort.TimSort(permutation)
	// fmt.Println(permutation)
	//	arrSz := 10000
	// arrSz := 10
	// mid := arrSz / 2
	// index := 0
	// array := make([]int, arrSz)
	// for i := 0; i < arrSz; i++ {
	// 	if i%2 == 0 {
	// 		array[index] = i
	// 		index++
	// 	} else {
	// 		array[mid] = i
	// 		mid++
	// 	}
	// }
	// fmt.Println(array)
	//

	matrix := make([][]int, 10)
	matrix[0] = append(matrix[0], 12, 11, 10, 12)
	fmt.Println(matrix)
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
	}
	fmt.Println(matrix)
}
