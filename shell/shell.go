// Package shell implements Shellsort.
package shell

import "cmp"

// Sort uses the Shellsort algorithm to sort a slice.
func Sort[T cmp.Ordered](s []T) {
	tokuda := [...]int{
		1147718700,
		510097200,
		226709866,
		100759940,
		44782196,
		19903198,
		8845866,
		3931496,
		1747331,
		776591,
		345152,
		153401,
		68178,
		30301,
		13467,
		5985,
		2660,
		1182,
		525,
		233,
		103,
		46,
		20,
		9,
		4,
		1,
	}

	for _, gap := range tokuda {
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
