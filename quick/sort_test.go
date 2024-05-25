package quick

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{
		15, 3, 2, 31, 19, 21, 6, 7,
		16, 4, 17, 30, 10, 11, 5, 18,
		25, 12, 8, 28, 14, 29, 22, 0,
		24, 13, 1, 20, 26, 27, 23, 9,
	}
	Sort(list)
	for i, v := range list {
		if i != v {
			t.Error(i, v)
		}
	}
}

func TestInsertion(t *testing.T) {
	list := []int{
		15, 3, 2, 31, 19, 21, 6, 7,
		16, 4, 17, 30, 10, 11, 5, 18,
		25, 12, 8, 28, 14, 29, 22, 0,
		24, 13, 1, 20, 26, 27, 23, 9,
	}
	insertion(list)
	for i, v := range list {
		if i != v {
			t.Error(i, v)
		}
	}
}

func TestSortFirst(t *testing.T) {
	list := []int{
		15, 3, 2, 31, 19, 21, 6, 7,
		16, 4, 17, 30, 10, 11, 5, 18,
		25, 12, 8, 28, 14, 29, 22, 0,
		24, 13, 1, 20, 26, 27, 23, 9,
	}
	SortFirst(list, 11)
	for i, v := range list[:11] {
		if i != v {
			t.Error(i, v)
		}
	}
}

func TestSelection(t *testing.T) {
	list := []int{
		15, 3, 2, 31, 19, 21, 6, 7,
		16, 4, 17, 30, 10, 11, 5, 18,
		25, 12, 8, 28, 14, 29, 22, 0,
		24, 13, 1, 20, 26, 27, 23, 9,
	}
	selection(list, 11)
	for i, v := range list[:11] {
		if i != v {
			t.Error(i, v)
		}
	}
}

func TestSelect(t *testing.T) {
	list := []int{
		15, 3, 2, 31, 19, 21, 6, 7,
		16, 4, 17, 30, 10, 11, 5, 18,
		25, 12, 8, 28, 14, 29, 22, 0,
		24, 13, 1, 20, 26, 27, 23, 9,
	}
	if k := Select(list, 3); k != 3 {
		t.Error(k)
	}
}

func BenchmarkSort(b *testing.B) {
	s := make([]float64, 10_000_000)
	for i := range s {
		s[i] = rand.Float64()
	}

	b.ResetTimer()
	Sort(s)
}

func BenchmarkSortK(b *testing.B) {
	s := make([]float64, 10_000_000)
	for i := range s {
		s[i] = rand.Float64()
	}

	b.ResetTimer()
	SortFirst(s, 1_000)
}

func BenchmarkSelect(b *testing.B) {
	s := make([]float64, 10_000_000)
	for i := range s {
		s[i] = rand.Float64()
	}

	b.ResetTimer()
	Select(s, 1_000_000)
}
