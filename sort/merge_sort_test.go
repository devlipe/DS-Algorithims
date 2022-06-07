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
