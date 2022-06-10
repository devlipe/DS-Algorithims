package sort_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/devlipe/data_structures/sort"
	"github.com/stretchr/testify/assert"
)

var inputs = []struct {
	input []int
}{
	{input: rand.Perm(100)},
	{input: rand.Perm(1000)},
	{input: rand.Perm(10000)},
	{input: rand.Perm(100000)},
	{input: rand.Perm(1000000)},
	{input: rand.Perm(10000000)},
}

func BenchmarkTimSort(b *testing.B) {
	for _, v := range inputs {
		b.Run(fmt.Sprintf("input_size_%d", len(v.input)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.TimSort(v.input)
			}
		})
	}
}

func TestTimSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	permutation := rand.Perm(1500)

	sort.MergeSort(permutation)
	assert.IsIncreasing(t, permutation, "O array deve estar ordenado")
}
