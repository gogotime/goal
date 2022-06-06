package tree

import (
	"fmt"
	"testing"
)

func TestSegTree(t *testing.T) {
	N := 1000000000
	book := [][]int{{22, 29}, {12, 17}, {20, 27}, {27, 36}, {24, 31}, {23, 28}, {47, 50}, {23, 30}, {24, 29}, {19, 25}, {19, 27}, {3, 9}, {34, 41}, {22, 27}, {3, 9}, {29, 38}, {34, 40}, {49, 50}, {42, 48}, {43, 50}, {39, 44}, {30, 38}, {42, 50}, {31, 39}, {9, 16}, {10, 18}, {31, 39}, {30, 39}, {48, 50}, {36, 42}}
	tree := NewSegTree[int](0, N, func(lv, rv int) int {
		return max(lv, rv)
	}, func(val, add int) int {
		return val + add
	})
	ans := []int{}
	for _, v := range book {
		start := v[0]
		end := v[1] - 1
		tree.Update(start, end, 1)
		ans = append(ans, tree.Query(0, N))
	}
	trueAns := []int{1, 1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	for i := range ans {
		if ans[i] != trueAns[i] {
			t.Fatal("not correct")
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
