package main

import (
	// "fmt"
	"math/rand"
	"time"

	"github.com/devlipe/data_structures/sort"
)

func main() {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1111150)

	// fmt.Println(permutation)
	sort.InsertionSort(permutation)
	//fmt.Println(permutation)

}
