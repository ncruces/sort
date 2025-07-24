// Package quick implements Tony Hoare's Quicksort and Quickselect.
//
// This package avoids quadratic behavior by using median-of-ninthers
// when a bad pivot is detected.
package quick

import "cmp"

const (
	minLen    = 32 // at least 1
	minK      = 4  // at least 1
	minMed3   = 32
	minRatio  = 16 // at least 1
	minMedNin = minRatio * 9
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

// SortLast uses the Quickselect and Quicksort algorithms to sort the last k elements of a slice.
// It uses O(n + k·log(k)) time and O(log(n)) space.
func SortLast[T cmp.Ordered](s []T, k int) {
	if k != 0 {
		n := len(s) - k
		Select(s, n)
		Sort(s[n+1:])
	}
}

// Select uses the Quickselect algorithm to find element k of the slice,
// partially sorting the slice around, and returning, s[k].
// It uses O(n) time and O(log₉(n)) space.
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
// use median-of-ninthers to select a better pivot.
// It uses O(n) time and O(log₉(n)) space.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1

	// For large r, sort 3 elements,
	// and use their median as a pivot.
	if r >= minMed3 {
		sort3(s, 0, r/2, r)
	}

	p := s[r/2]
	i := hoarePartition(s, p)

	// For really large r, check if the pivot was bad,
	// and use median-of-ninthers to pick a better one.
	if r >= minMedNin {
		b := r / minRatio
		if !(b < i && i < r-b) {
			p = medianOfNinthers(s)
			i = hoarePartition(s, p)
		}
	}
	return i
}

// HoarePartition implements Hoare's partition scheme (not Lomuto).
// Hoare's partition handles repeated elements sensibly.
// It uses O(n) time and O(1) space.
func hoarePartition[T cmp.Ordered](s []T, p T) int {
	i := 0
	j := len(s) - 1
	for {
		for cmp.Less(s[i], p) {
			i += 1
		}
		for cmp.Less(p, s[j]) {
			j -= 1
		}
		if i >= j {
			return j + 1
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
		for j, q := range s[i+1:] {
			if cmp.Less(q, p) {
				m = j + 1
				p = q
			}
		}
		s[i], s[m+i] = s[m+i], s[i]
	}
}

// MedianOfNinthers fills the middle ninth of the slice
// with the ninthers of 9-tuples taken from the slice,
// then uses Quickselect to find their median.
// It uses O(n) time and O(log₉(n)) space.
func medianOfNinthers[T cmp.Ordered](s []T) T {
	s = mediansOfTriples(s)
	s = mediansOfTriples(s)
	return Select(s, len(s)/2)
}

// MediansOfTriples returns the middle third of the slice
// filled with the medians of triples taken from the slice.
// It uses O(n) time and O(1) space.
func mediansOfTriples[T cmp.Ordered](s []T) []T {
	n := len(s) / 3
	for i := range n {
		sort3(s, i, i+n, i+n+n)
	}
	return s[n : n+n]
}

// Sort3 sorts three elements of the slice.
func sort3[T cmp.Ordered](s []T, i, j, k int) {
	if cmp.Less(s[j], s[i]) {
		s[i], s[j] = s[j], s[i]
	}
	if cmp.Less(s[k], s[j]) {
		s[j], s[k] = s[k], s[j]
		if cmp.Less(s[j], s[i]) {
			s[i], s[j] = s[j], s[i]
		}
	}
}
