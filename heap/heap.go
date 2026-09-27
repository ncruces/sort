// Package heap implements Floyd's bottom-up Heapsort.
package heap

import "cmp"

// Sort uses the Heapsort algorithm to sort a slice.
// It uses O(n·log(n)) time and O(1) space.
func Sort[T cmp.Ordered](s []T) {
	heapify(s)
	m := len(s)
	for {
		if m--; m <= 0 {
			break
		}
		s[0], s[m] = s[m], s[0]
		siftDown(s[:m], 0)
	}
}

// SortLast uses the Heapsort algorithm to sort the last k elements of a slice.
// It uses O(n + k·log(n)) time and O(1) space.
func SortLast[T cmp.Ordered](s []T, k int) {
	// This does a bounds check before making any changes to the slice.
	_ = s[:k:len(s)]

	heapify(s)
	m := len(s)
	for range k {
		if m--; m <= 0 {
			break
		}
		s[0], s[m] = s[m], s[0]
		siftDown(s[:m], 0)
	}
}

// SortFirst uses the Heapselect and Heapsort algorithms to sort the first k elements of a slice.
// It uses O(k + (n-k)·log(k)) time and O(1) space.
func SortFirst[T cmp.Ordered](s []T, k int) {
	// This does a bounds check before making any changes to the slice.
	_ = s[:k:len(s)]

	heapify(s[:k])
	for m := k; m < len(s); m++ {
		if cmp.Less(s[m], s[0]) {
			s[0], s[m] = s[m], s[0]
			siftDown(s[:k], 0)
		}
	}
	for {
		if k--; k <= 0 {
			break
		}
		s[0], s[k] = s[k], s[0]
		siftDown(s[:k], 0)
	}
}

// Select uses the Heapselect algorithm to find element k of the slice,
// partially sorting the slice around, and returning, s[k].
// It uses O(k + (n-k)·log(k)) time and O(1) space.
func Select[T cmp.Ordered](s []T, k int) T {
	// This does a bounds check before making any changes to the slice.
	_ = s[k]

	n := k + 1
	heapify(s[:n])
	for m := n; m < len(s); m++ {
		if cmp.Less(s[m], s[0]) {
			s[0], s[m] = s[m], s[0]
			siftDown(s[:n], 0)
		}
	}
	s[0], s[k] = s[k], s[0]
	return s[k]
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
