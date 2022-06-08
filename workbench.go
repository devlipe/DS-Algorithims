package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/devlipe/data_structures/sort"
)

func main() {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(10)

	fmt.Println(permutation)
	sort.QuickSort(permutation)
	fmt.Println(permutation)
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
}
