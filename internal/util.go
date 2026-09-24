package internal

import (
	"iter"
	"math/bits"
	"math/rand"
	"slices"
)

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
	r := rand.New(rand.NewSource(42))
	return r.Perm(n)
}

func enum(n int) []int {
	s := permutation(n)
	for i := range s {
		s[i] &= 3
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
	for i := range m {
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

func Ints(n int) iter.Seq2[string, []int] {
	return func(yield func(string, []int) bool) {
		if true &&
			yield("zeros", zeros(n)) &&
			yield("enums", enum(n)) &&
			yield("sorted", sorted(n)) &&
			yield("rotated", rotated(n)) &&
			yield("reversed", reversed(n)) &&
			yield("pipeorgan", pipeorgan(n)) &&
			yield("permutation", permutation(n)) &&
			yield("killer", killer(1<<bits.Len(uint(n))-1)) {
		}
	}
}

func Floats(n int) []float64 {
	r := rand.New(rand.NewSource(42))
	s := make([]float64, n)
	for i := range s {
		s[i] = r.Float64()
	}
	return s
}
