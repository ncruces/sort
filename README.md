# Sorting algorithms

[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/ncruces/sort)
[![Go Report](https://goreportcard.com/badge/github.com/ncruces/sort)](https://goreportcard.com/report/github.com/ncruces/sort)
[![Go Coverage](https://github.com/ncruces/sort/wiki/coverage.svg)](https://raw.githack.com/wiki/ncruces/sort/coverage.html)

Sorting algorithms implemented in Go.\
Use as a learning resource.\
Clarity and simplicity, not so much performance, are the goals.

## Quicksort

The most interesting (and more practical) algorithm here is Tony Hoare's Quicksort.
It's relatively fast, only slightly slower than the standard library on the average case.
It's flexible, and gives you asymptotically optimal algorithms for median finding, top-K, etc
(which the standard library lacks).

![quicksort visualization](anims/quick.png)

This version avoids quadratic behavior by using median-of-ninthers when a bad pivot is detected:

![median-of-ninthers visualization](anims/ninthers.png)

The algorithm is fully deterministic,
and every step contributes to partially sorting the array.

## Heapsort

An implementation of Floyd's bottom-up Heapsort.

![heapsort visualization](anims/heap.png)

## Shellsort

An implementation of Shellsort with the Gonnet & Baeza-Yates gap sequence.

![heapsort visualization](anims/shell.png)

## Credits

Visualizations thanks to:
[`github.com/invzhi/sorting-visualization`](https://github.com/invzhi/sorting-visualization/)