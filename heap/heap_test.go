package heap

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(1_000_000)},
		{"bits", bits(1_000_000)},
		{"sorted", sorted(1_000_000)},
		{"rotated", rotated(1_000_000)},
		{"reversed", reversed(1_000_000)},
		{"pipeorgan", pipeorgan(1_000_000)},
		{"permutation", permutation(1_000_000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.list)
			if !slices.IsSorted(tt.list) {
				t.FailNow()
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	list := floats(10_000_000)
	b.ResetTimer()
	Sort(list)
}

func zeros(n int) []int {
	return make([]int, n)
}

func sorted(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

func rotated(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i + 1
	}
	s[n-1] = 0
	return s
}

func reversed(n int) []int {
	s := sorted(n)
	slices.Reverse(s)
	return s
}

func permutation(n int) []int {
	return rand.Perm(n)
}

func bits(n int) []int {
	s := rand.Perm(n)
	for i := range s {
		s[i] &= 1
	}
	return s
}

func floats(n int) []float64 {
	s := make([]float64, n)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}

func pipeorgan(n int) []int {
	return append(sorted(n/2), reversed(n/2)...)
}
