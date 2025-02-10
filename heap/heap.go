// Package heap implements Floyd's bottom-up Heapsort.
package heap

import "cmp"

// Sort uses the Heapsort algorithm to sort a slice.
// It uses O(n·log(n)) time and O(1) space.
func Sort[T cmp.Ordered](s []T) {
	heapify(s)

	m := len(s)
	for m > 1 {
		m -= 1
		s[0], s[m] = s[m], s[0]
		siftDown(s[:m], 0)
	}
}

// Heapify rearranges a slice into a binary max-heap.
// It uses O(n) time and O(1) space.
func heapify[T cmp.Ordered](s []T) {
	for i := len(s)/2 - 1; i >= 0; i -= 1 {
		siftDown(s, i)
	}
}

// SiftDown is the core of the Heapsort algorithm.
// It constructs binary heaps out of smaller heaps.
// It uses O(log(n)) time and O(1) space.
func siftDown[T cmp.Ordered](s []T, i int) {
	t := s[i]
	j := minSearch(s, i)
	for cmp.Less(s[j], t) {
		j = (j - 1) / 2
	}
	for j > i {
		s[j], t = t, s[j]
		j = (j - 1) / 2
	}
	s[i] = t
}

// MinSearch searches for the leaf where
// the minimum possible value would be placed.
// It uses O(log(n)) time and O(1) space.
func minSearch[T cmp.Ordered](s []T, j int) int {
	for {
		l := 2*j + 1
		r := 2*j + 2
		switch {
		case r > len(s):
			return j
		case r == len(s):
			return l
		}
		if cmp.Less(s[l], s[r]) {
			j = r
		} else {
			j = l
		}
	}
}
