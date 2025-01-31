// Package quick implements Tony Hoare's Quicksort and Quickselect.
//
// This package avoids quadratic behavior by using median-of-ninthers
// when a bad pivot is detected.
package quick

import "cmp"

const (
	minLen    = 32 // at least 1
	minK      = 4  // at least 1
	minMed3   = 32 // at least 1
	minRatio  = 16 // at least 4
	minMedNin = 144
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
// use median-of-ninthers to select a better pivot.
// It uses O(n) time and O(log(n)) space.
func partition[T cmp.Ordered](s []T) int {
	r := len(s) - 1

	// For large r, sort 3 elements,
	// and use their median as a pivot.
	if r >= minMed3 {
		if cmp.Less(s[r/2], s[0]) {
			s[0], s[r/2] = s[r/2], s[0]
		}
		if cmp.Less(s[r], s[r/2]) {
			s[r], s[r/2] = s[r/2], s[r]
			if cmp.Less(s[r/2], s[0]) {
				s[0], s[r/2] = s[r/2], s[0]
			}
		}
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

// MedianOfNinthers fills the first ninth of the slice
// with the ninthers of all 9-tuples of the slice,
// then uses Quickselect to find their median.
// It uses O(n) time and O(log(n)) space.
func medianOfNinthers[T cmp.Ordered](s []T) T {
	m := mediansOfTriples(s)
	n := mediansOfTriples(s[:m])
	if n < 2 {
		return s[0]
	}
	return Select(s[:n], n/2)
}

// MediansOfTriples fills the first third of the slice
// with the medians of all triples of the slice.
func mediansOfTriples[T cmp.Ordered](s []T) (m int) {
	for i := 0; i+3 <= len(s); i += 3 {
		t := s[i : i+3]
		// Sort 3 elements.
		if cmp.Less(t[1], t[0]) {
			t[0], t[1] = t[1], t[0]
		}
		if cmp.Less(t[2], t[1]) {
			t[1], t[2] = t[2], t[1]
			if cmp.Less(t[1], t[0]) {
				t[0], t[1] = t[1], t[0]
			}
		}
		s[m], t[1] = t[1], s[m]
		m += 1
	}
	return m
}
