// Package quick implements Tony Hoare's Quicksort and Quickselect.
//
// This package avoids quadratic behavior by using Median-of-medians
// when a bad pivot is detected.
package quick

import "cmp"

const (
	minLen    = 32 // at least 1
	minK      = 4  // at least 1
	minMed3   = 32 // at least 1
	minRatio  = 16 // at least 4
	minMedMed = 128
)

// Sort uses the Quicksort algorithm to sort a slice.
// It uses O(n·log(n)) time and O(log(n)) space.
func Sort[T cmp.Ordered](s []T) {
	// We could check for len(s) > 1, and use Quicksort all the way down.
	// In practise, Insertion sort performs better at small sizes.
	for len(s) > minLen {
		p := partition(s)
		// Recursing into the smaller side conserves stack space.
		if p > len(s)/2 {
			Sort(s[p:])
			s = s[:p]
		} else {
			Sort(s[:p])
			s = s[p:]
		}
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
// This bit only does pivot selection:
// - the middle element for small slices,
// - the median of 3 for bigger slices.
// This produces optimal partitions in many common cases.
// If it turns out to be a really bad choice,
// use Median-of-medians to select a better pivot.
// It uses O(n) time and O(log(n)) space.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1

	// For large r, sort 3 elements,
	// and use their median as a pivot.
	if r >= minMed3 {
		if cmp.Less(s[r], s[0]) {
			s[0], s[r] = s[r], s[0]
		}
		if cmp.Less(s[r/2], s[0]) {
			s[0], s[r/2] = s[r/2], s[0]
		}
		if cmp.Less(s[r], s[r/2]) {
			s[r], s[r/2] = s[r/2], s[r]
		}
	}

	p := s[r/2]
	i := hoarePartition(s, p)

	// For really large r, check if the pivot was bad,
	// and use Median-of-medians to pick a better one.
	if r >= minMedMed {
		b := r / minRatio
		if !(b < i && i < r-b) {
			p = medianOfMedians(s)
			i = hoarePartition(s, p)
		}
	}
	return i
}

// HoarePartition implements Hoare's partition scheme (not Lomuto).
// Hoare's partition handles repeated elements sensibly.
func hoarePartition[T cmp.Ordered](s []T, p T) int {
	r := len(s) - 1
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

// MedianOfMedians selects a good pivot for partition.
// A good pivot lies in the middle 40% of the slice.
// It uses O(n) time and O(log(n)) space.
func medianOfMedians[T cmp.Ordered](s []T) T {
	m := 0
	for i := 0; i+5 < len(s); i += 5 {
		// Sort groups of 5 elements and move their medians
		// to the start of the slice.
		insertion(s[i : i+5])
		s[m], s[i+2] = s[i+2], s[m]
		m += 1
	}
	if m < 2 {
		return s[0]
	}
	// Use Quickselect to find the Median-of-medians.
	return Select(s[:m], m/2)
}
