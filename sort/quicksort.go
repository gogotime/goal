package sort

import (
	"goal/constraint"
	"math/rand"
)

func partition[T constraint.Comparable](a []T) int {
	l, r := 0, len(a)-1
	x := a[r]
	i := -1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func partitionDesc[T constraint.Comparable](a []T) int {
	l, r := 0, len(a)-1
	x := a[r]
	i := -1
	for j := l; j < r; j++ {
		if a[j] >= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func randomPartition[T constraint.Comparable](a []T) int {
	if len(a) <= 1 {
		return 0
	}
	r := len(a) - 1
	i := rand.Intn(r)
	a[i], a[r] = a[r], a[i]
	return partition(a)
}

func randomPartitionDesc[T constraint.Comparable](a []T) int {
	if len(a) <= 1 {
		return 0
	}
	r := len(a) - 1
	i := rand.Intn(r)
	a[i], a[r] = a[r], a[i]
	return partitionDesc(a)
}

func QuickSelect[T constraint.Comparable](a []T, index int) T {
	if a == nil {
		return *new(T)
	}
	if index < 0 {
		return a[0]
	}
	q := randomPartition(a)
	if q == index {
		return a[q]
	}
	if q < index {
		return QuickSelect(a[q+1:], index-q-1)
	}
	return QuickSelect(a[:q], index)
}

func QuickSort[T constraint.Comparable](a []T) {
	if a == nil || len(a) <= 1 {
		return
	}
	q := randomPartition(a)
	QuickSort(a[:q])
	QuickSort(a[q+1:])
}

func QuickSortDesc[T constraint.Comparable](a []T) {
	if a == nil || len(a) <= 1 {
		return
	}
	q := randomPartitionDesc(a)
	QuickSortDesc(a[:q])
	QuickSortDesc(a[q+1:])
}
