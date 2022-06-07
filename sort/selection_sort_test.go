package sort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/devlipe/data_structures/sort"
	"github.com/stretchr/testify/assert"
)

///Function to test if selection sort is working
func TestSelectionSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(2000)

	sort.SelectionSort(permutation)
	assert.IsIncreasing(t, permutation, "O array deve estar ordenado")
}
