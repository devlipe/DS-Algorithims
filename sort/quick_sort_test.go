package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/devlipe/data_structures/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1500)

	sort.MergeSort(permutation)
	assert.IsIncreasing(t, permutation, "O array deve estar ordenado")
}

func TestPartition(t *testing.T) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1500)

	pos := sort.Partition(permutation, 0, len(permutation))

	for _, v := range permutation[:pos] {
		assert.Less(t, v, permutation[pos], "All the elements before pivot: %d must be less than it. found: %v")
	}
	for _, v := range permutation[pos+1:] {
		assert.Greaterf(t, v, permutation[pos], "All the elements after pivot: %d must be greater than it. found: %v", permutation[pos], v)

	}
}
