package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/devlipe/data_structures/sort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1500)

	sort.MergeSort(permutation)
	assert.IsIncreasing(t, permutation, "O array deve estar ordenado")
}

func TestMerge(t *testing.T) {
	arrSz := 10000
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

	sort.Merge(array, 0, arrSz/2, arrSz)
	assert.IsIncreasing(t, array, "The array must be sorted correctly")
}
