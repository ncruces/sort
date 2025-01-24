package quick

import (
	"cmp"
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
		{"reversed", reversed(1_000_000)},
		{"pipeorgan", pipeorgan(1_000_000)},
		{"permutation", permutation(1_000_000)},
		{"killer", killer(1024*1024 - 1)},
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

func TestSortFirst(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(1_000_000)},
		{"bits", bits(1_000_000)},
		{"sorted", sorted(1_000_000)},
		{"reversed", reversed(1_000_000)},
		{"pipeorgan", pipeorgan(1_000_000)},
		{"permutation", permutation(1_000_000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortFirst(tt.list, 1111)
			if !slices.IsSorted(tt.list[:1111]) {
				t.FailNow()
			}
		})
	}
}

func TestSortLast(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(1_000_000)},
		{"bits", bits(1_000_000)},
		{"sorted", sorted(1_000_000)},
		{"reversed", reversed(1_000_000)},
		{"pipeorgan", pipeorgan(1_000_000)},
		{"permutation", permutation(1_000_000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortLast(tt.list, 1111)
			n := len(tt.list) - 1111
			if !slices.IsSorted(tt.list[n:]) {
				t.FailNow()
			}
		})
	}
}

func TestSelect(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(1_000_000)},
		{"bits", bits(1_000_000)},
		{"sorted", sorted(1_000_000)},
		{"reversed", reversed(1_000_000)},
		{"pipeorgan", pipeorgan(1_000_000)},
		{"permutation", permutation(1_000_000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sel := Select(tt.list, 1111)
			slices.Sort(tt.list)
			if sel != tt.list[1111] {
				t.FailNow()
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(100)},
		{"bits", bits(100)},
		{"sorted", sorted(100)},
		{"reversed", reversed(100)},
		{"pipeorgan", pipeorgan(100)},
		{"permutation", permutation(100)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertion(tt.list)
			if !slices.IsSorted(tt.list) {
				t.FailNow()
			}
		})
	}
}

func TestSelection(t *testing.T) {
	tests := []struct {
		name string
		list []int
	}{
		{"zeros", zeros(100)},
		{"bits", bits(100)},
		{"sorted", sorted(100)},
		{"reversed", reversed(100)},
		{"pipeorgan", pipeorgan(100)},
		{"permutation", permutation(100)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selection(tt.list, 11)
			if !slices.IsSorted(tt.list[:11]) {
				t.FailNow()
			}
		})
	}
}

func TestBounds(t *testing.T) {
	Sort[int](nil)
	Sort([]int{0})

	SortFirst[int](nil, 0)
	SortFirst([]int{0}, 1)

	SortLast[int](nil, 0)
	SortLast([]int{0}, 1)

	Select([]int{0}, 0)

	partition([]int{0})
	insertion[int](nil)
	selection[int](nil, 0)
	medianOfMedians([]int{0})
}

func FuzzPartition(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		if len(s) < 2 {
			t.SkipNow()
		}

		i := partition(s)

		if len(s[:i]) == 0 || len(s[i:]) == 0 {
			t.FailNow()
		}
		if cmp.Less(slices.Min(s[i:]), slices.Max(s[:i])) {
			t.FailNow()
		}
	})
}

func BenchmarkSort(b *testing.B) {
	list := floats(10_000_000)
	b.ResetTimer()
	Sort(list)
}

func BenchmarkSortFirst(b *testing.B) {
	list := floats(10_000_000)
	b.ResetTimer()
	SortFirst(list, 10_000)
}

func BenchmarkSortLast(b *testing.B) {
	list := floats(10_000_000)
	b.ResetTimer()
	SortLast(list, 10_000)
}

func BenchmarkSelect(b *testing.B) {
	list := floats(10_000_000)
	b.ResetTimer()
	Select(list, 1_000_000)
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

func killer(n int) []int {
	// https://webpages.charlotte.edu/rbunescu/courses/ou/cs4040/introsort.pdf

	s := make([]int, n)

	if n%2 != 0 {
		s[n-1] = n
		n--
	}

	m := n / 2
	for i := 0; i < m; i++ {
		// first half of array
		if i%2 == 0 {
			// even indices
			s[i] = i + 1
		} else {
			// odd indices
			s[i] = i + m + (m & 1)
		}
		// second half of array
		s[m+i] = (i + 1) * 2
	}

	return s
}
