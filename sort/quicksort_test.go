package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	arr := []int{}
	QuickSort(arr)

	arr = []int{1}
	QuickSort(arr)
	if arr[0] != 1 {
		t.Fatal("QuickSort not correct when len(arr)==1")
	}

	arr = []int{3, 1}
	QuickSort(arr)
	if arr[0] != 1 && arr[1] != 3 {
		t.Fatal("QuickSort not correct when len(arr)==2")
	}

	rand.Seed(time.Now().Unix())
	n := 200000
	arr = make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int())
	}
	QuickSort(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			t.Fatal("QuickSort not correct when len(arr)==n", arr[i-1], "<", arr[i])
		}
	}
}

func TestQuickSortDesc(t *testing.T) {
	arr := []int{}
	QuickSortDesc(arr)

	arr = []int{1}
	QuickSortDesc(arr)
	if arr[0] != 1 {
		t.Fatal("QuickSortDesc not correct when len(arr)==1")
	}

	arr = []int{1, 3}
	QuickSortDesc(arr)
	if arr[0] != 3 && arr[1] != 1 {
		t.Fatal("QuickSortDesc not correct when len(arr)==2")
	}

	rand.Seed(time.Now().Unix())
	n := 200000
	arr = make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int())
	}
	QuickSortDesc(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			t.Fatal("QuickSortDesc not correct when len(arr)==n")
		}
	}
}
