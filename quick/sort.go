// Package quick implements Tony Hoare's Quicksort and Quickselect.
//
// This package avoids quadratic behavior by using Median-of-medians
// when a bad pivot is detected.
package quick

import "cmp"

const (
	minLen    = 32 // at least 1
	minK      = 4  // at least 1
	minRatio  = 16 // at least 4
	minMedian = 20 // at least 20
)

// Sort uses the Quicksort algorithm to sort a slice.
// It uses O(n·log(n)) time and O(log(n)) space.
func Sort[T cmp.Ordered](s []T) {
	// We could check for len(s) > 1, and use Quicksort all the way down.
	// In practise, Insertion sort performs better at small sizes.
	for len(s) > minLen {
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

	// We could check for len(s) > 1, and use Quickselect all the way down.
	// In practise, Selection sort performs better for small k.
	for k > minK {
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

	// We could check for len(s) > 1, and use Quickselect all the way down.
	// In practise, Selection sort performs better for small k.
	for k >= minK {
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

// Partition is the core of the Quicksort and Quickselect algorithms.
// This version uses the middle element as a pivot,
// producing an optimal partition in many common cases.
// If this turns out to be a terrible choice,
// Median-of-medians is used to select a good pivot.
// It uses O(n) time and O(log(n)) space.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1
	p := s[r/2]
retry:
	i := 0
	j := r
	for {
		for i < r && cmp.Less(s[i], p) {
			i += 1
		}
		for j > 0 && cmp.Less(p, s[j]) {
			j -= 1
		}
		if i > j {
			if badPartition(i, r) {
				p = medianOfMedians(s)
				goto retry
			}
			return i
		}
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}

// Insertion sort is used as the base case for Quicksort.
// It uses O(n²) time and O(1) space (used for small n).
func insertion[T cmp.Ordered](s []T) {
	for i, p := range s {
		for i > 0 && cmp.Less(p, s[i-1]) {
			s[i] = s[i-1]
			i -= 1
		}
		s[i] = p
	}
}

// Selection sort is used as the base case for Quickselect.
// It uses O(n·k) time and O(1) space (used for small k).
func selection[T cmp.Ordered](s []T, k int) {
	for i, p := range s[:k] {
		m := 0
		for j, q := range s[i:] {
			if cmp.Less(q, p) {
				m = j
				p = q
			}
		}
		s[i], s[m+i] = s[m+i], s[i]
	}
}

// BadPartition identifies terrible partitions.
// To ensure termination, a bad pivot must lie
// outside the middle 50% of the slice.
func badPartition(i, r int) bool {
	b := r / minRatio
	return b > minMedian && !(b < i && i < r-b)
}

// MedianOfMedians selects a good pivot for partition.
// A good pivot lies in the middle 40% of the slice.
// It uses O(n) time and O(log(n)) space.
func medianOfMedians[T cmp.Ordered](s []T) T {
	m := 0
	for i := 0; i+5 < len(s); i += 5 {
		insertion(s[i : i+5])
		s[m], s[i+2] = s[i+2], s[m]
		m += 1
	}
	if m < 2 {
		return s[0]
	}
	return Select(s[:m], m/2)
}
