package sort

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
	slices.Sort(Floats(1_000_000))
}

func BenchmarkStandard(b *testing.B) {
	list := Floats(10_000_000)
	b.Run("floats", func(b *testing.B) {
		slices.Sort(slices.Clone(list))
	})
	for name, list := range Ints(10_000_000) {
		b.Run(name, func(b *testing.B) {
			slices.Sort(slices.Clone(list))
		})
	}
}
