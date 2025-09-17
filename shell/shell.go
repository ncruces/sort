// Package shell implements Shellsort with the
// Gonnet & Baeza-Yates gap sequence.
package shell

import "cmp"

// Sort uses the Shellsort algorithm to sort a slice.
func Sort[T cmp.Ordered](s []T) {
	for gap := len(s); gap > 1; {
		gap = int(max(1, (uint64(gap)*5-1)/11))
		for i := gap; i < len(s); i += 1 {
			j, p := i, s[i]
			for j >= gap && cmp.Less(p, s[j-gap]) {
				s[j] = s[j-gap]
				j -= gap
			}
			s[j] = p
		}
	}
}
