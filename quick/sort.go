// Package quick implements Tony Hoare's Quicksort and Quickselect.
package quick

import "cmp"

// Sort uses the Quicksort algorithm to sort a slice.
// It uses O(n·log(n)) time and O(log(n)) space.
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
// It uses O(n + k·log(k)) time and O(log(n)) space.
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

// Select uses the Quickselect algorithm to find element k of the slice,
// partially sorting the slice around, and returning, s[k].
// It uses O(n) time and O(log(n)) space.
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

// Partition is the core of the Quicksort algorithm.
// This version uses the middle element as a pivot,
// producing an optimal partition in many common cases.
// If this turns out to be a terrible choice,
// Median-of-medians is used to select a good pivot.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1
	p := s[r/2]
retry:
	i := 0
	j := r
	for {
		for i < r && cmp.Less(s[i], p) {
			i++
		}
		for j > 0 && cmp.Less(p, s[j]) {
			j--
		}
		if i > j {
			if badPartition(i, r) {
				p = medianOfMedians(s)
				goto retry
			}
			return i
		}
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
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

// BadPartition identifies terrible partitions.
func badPartition(i, r int) bool {
	b := r / 16
	return b > 4 && !(b < i && i < r-b)
}

// MedianOfMedians selects a good pivot for partition.
func medianOfMedians[T cmp.Ordered](s []T) T {
	m := 0
	for i := 0; i+5 < len(s); i += 5 {
		insertion(s[i : i+5])
		s[m], s[i+2] = s[i+2], s[m]
		m++
	}
	return Select(s[:m], m/2)
}
