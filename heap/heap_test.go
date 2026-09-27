package heap

import (
	"flag"
	"os"
	"slices"
	"testing"

	. "github.com/ncruces/sort/internal"
)

func TestMain(m *testing.M) {
	_ = flag.Set("test.benchtime", "1x")
	os.Exit(m.Run())
}

func TestSort(t *testing.T) {
	for name, list := range Ints(1_000_000) {
		t.Run(name, func(t *testing.T) {
			Sort(list)
			if !slices.IsSorted(list) {
				t.FailNow()
			}
		})
	}
}

func TestSortFirst(t *testing.T) {
	for name, list := range Ints(1_000_000) {
		t.Run(name, func(t *testing.T) {
			SortFirst(list, 1111)
			if !slices.IsSorted(list[:1111]) {
				t.FailNow()
			}
		})
	}
}

func TestSortLast(t *testing.T) {
	for name, list := range Ints(1_000_000) {
		t.Run(name, func(t *testing.T) {
			SortLast(list, 1111)
			n := len(list) - 1112
			if !slices.IsSorted(list[n:]) {
				t.FailNow()
			}
		})
	}
}

func TestSelect(t *testing.T) {
	for name, list := range Ints(1_000_000) {
		t.Run(name, func(t *testing.T) {
			sel := Select(list, 1111)
			if sel != list[1111] {
				t.FailNow()
			}
			slices.Sort(list[:1111])
			slices.Sort(list[1112:])
			if !slices.IsSorted(list) {
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
}

func BenchmarkSort(b *testing.B) {
	list := Floats(10_000_000)
	b.Run("floats", func(b *testing.B) {
		Sort(slices.Clone(list))
	})
	for name, list := range Ints(10_000_000) {
		b.Run(name, func(b *testing.B) {
			Sort(slices.Clone(list))
		})
	}
}

func BenchmarkSortFirst(b *testing.B) {
	list := Floats(10_000_000)
	b.ResetTimer()
	SortFirst(list, 10_000)
}

func BenchmarkSortLast(b *testing.B) {
	list := Floats(10_000_000)
	b.ResetTimer()
	SortLast(list, 10_000)
}

func BenchmarkSelect(b *testing.B) {
	list := Floats(10_000_000)
	b.ResetTimer()
	Select(list, 1_000_000)
}
