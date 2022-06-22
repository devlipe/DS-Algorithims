package sort_test

import (
	"testing"

	"github.com/devlipe/data_structures/sort"
)

func TestBinarySearch(t *testing.T) {
	a := []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}

	inputs := []struct {
		pos int
	}{
		{0},
		{58},
		{999},
		{-1},
		{480},
		{750},
		{666},
		{333},
		{150},
		{15}}

	t.Run("BinarySearch", func(t *testing.T) {
		for _, input := range inputs {
			if pos := sort.BinarySearch(a, input.pos); pos != input.pos {
				t.Errorf("BinarySearch(%v) = %v, want %v", input.pos, pos, input.pos)
			}
		}
	})
}

func TestBinarySearchInteractive(t *testing.T) {
	a := []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}

	inputs := []struct {
		pos int
	}{
		{0},
		{58},
		{999},
		{-1},
		{480},
		{750},
		{666},
		{333},
		{150},
		{15}}

	t.Run("BinarySearch", func(t *testing.T) {
		for _, input := range inputs {
			if pos := sort.BinarySearchInterative(a, input.pos); pos != input.pos {
				t.Errorf("BinarySearch(%v) = %v, want %v", input.pos, pos, input.pos)
			}
		}
	})
}

func BenchmarkBinarySearch(b *testing.B) {
	a := []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}
	inputs := []struct {
		pos int
	}{
		{0},
		{58},
		{999},
		{-1},
		{480},
		{750},
		{666},
		{333},
		{150},
		{15}}
	for _, input := range inputs {
		b.Run("BinarySearch", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.BinarySearch(a, input.pos)
			}
		})
	}
}

func BenchmarkSequentialSearch(b *testing.B) {
	a := []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}
	inputs := []struct {
		pos int
	}{
		{0},
		{58},
		{999},
		{-1},
		{480},
		{750},
		{666},
		{333},
		{150},
		{15}}
	for _, input := range inputs {
		b.Run("SequentialSearch", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SequentialSearch(a, input.pos)
			}
		})
	}
}
