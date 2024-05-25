package quick

import "cmp"

// Sort uses the Quicksort algorithm to sort a slice.
// On average the algorithm requres O(n·log(n)) space and time,
// but the worse case is O(n²).
func Sort[T cmp.Ordered](s []T) {
	// We could check for len(s) > 1, and use Quicksort all the way down.
	// In practise, Insertion sort performs better at small sizes.
	for len(s) > 32 {
		p := partition(s)
		Sort(s[:p])
		s = s[p:]
	}
	insertion(s)
}

// SortFirst uses the Quickselect and Quicksort algorithms to sort the first k elements of a slice.
// On average the algorithm requres O(n + k·log(k)) space and time,
// but the worse case is O(n²).
func SortFirst[T cmp.Ordered](s []T, k int) {
	// This does a bounds check before making any changes to the slice.
	_ = s[:k]

	// We could check for k > 0, and use Quickselect all the way down.
	// In practise, Selection sort performs better at small sizes.
	for k > 4 {
		p := partition(s)
		if p > k {
			s = s[:p]
		} else {
			Sort(s[:p])
			s = s[p:]
			k -= p
		}
	}
	selection(s, k)
}

// Select uses the Quickselect algorithm to find element k of the slice, if sorted.
// It partially sorts the slice around, and returns, s[k].
// On average the algorithm requres O(n) space and time,
// but the worse case is O(n²).
func Select[T cmp.Ordered](s []T, k int) T {
	// This does a bounds check before making any changes to the slice.
	_ = s[k]

	// We could check for k > 0, and use Quickselect all the way down.
	// In practise, Selection sort performs better at small sizes.
	for k >= 4 {
		p := partition(s)
		if p > k {
			s = s[:p]
		} else {
			s = s[p:]
			k -= p
		}
	}
	selection(s, k+1)
	return s[k]
}

// Partition is the core of any Quicksort algorithm.
// This version uses the middle element as a pivot.
// It avoids the pitfalls of an already sorted,
// reverse sorted, or all equal elements slice
// being the worse cases for partition.
// With this algorithm, those rather common cases
// produce perfectly balanced partions instead.
// In practise, this means that repeatedly using
// Select and SortFirst on the same slice
// does not exhibit worse case performance.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1
	p := s[r/2]
	i := 0
	j := r
	for i <= j {
		for i < r && cmp.Less(s[i], p) {
			i++
		}
		for j > 0 && cmp.Less(p, s[j]) {
			j--
		}
		if i <= j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	return i
}

// Insertion sort is used as the base case for Quicksort.
func insertion[T cmp.Ordered](s []T) {
	for i, p := range s {
		for i > 0 && cmp.Less(p, s[i-1]) {
			s[i] = s[i-1]
			i--
		}
		s[i] = p
	}
}

// Selection sort is used as the base case for Quickselect.
func selection[T cmp.Ordered](s []T, k int) {
	for i, p := range s[:k] {
		m := i
		for j, q := range s[i:] {
			if cmp.Less(q, p) {
				m = i + j
				p = q
			}
		}
		s[i], s[m] = s[m], s[i]
	}
}
