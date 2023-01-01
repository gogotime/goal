package list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickList_Insert(t *testing.T) {
	ql := newRandQuickList(1000)

	checkQuickListDistance(t, ql)
}

func TestQuickList_FirstKey(t *testing.T) {
	ql := newOrderedQuickList(100000)

	key := ql.FirstKey()

	if key != 0 {
		t.Errorf("failed")
	}
}

func TestQuickList_Last(t *testing.T) {
	ql := newOrderedQuickList(100000)

	key := ql.LastKey()

	if key != 100000-1 {
		t.Errorf("failed got key=%d, expect=%d", key, 100000-1)
	}
}

func TestQuickList_Rank(t *testing.T) {
	ql := newOrderedQuickList(100000)

	rank := ql.Rank(1234)
	if rank != 1234 {
		t.Errorf("failed got rank=%d, expect=%d", rank, 1234)
	}
}

func TestQuickList_NextKey(t *testing.T) {
	ql := newOrderedQuickList(100000)

	nk := ql.NextKey(99999)
	if nk != -1 {
		t.Errorf("failed got rank=%d, expect=%d", nk, -1)
	}

	nk = ql.NextKey(99998)
	if nk != 99999 {
		t.Errorf("failed got rank=%d, expect=%d", nk, 99999)
	}
}

func newRandQuickList(n int) *QuickList {
	sl := NewQuickList()
	for i := 0; i < n; i++ {
		sl.Insert(i, rand.Int()%n)
	}
	return sl
}

func newOrderedQuickList(n int) *QuickList {
	sl := NewQuickList()
	for i := 0; i < n; i++ {
		sl.Insert(i, i)
	}
	return sl
}

func checkQuickListDistance(t *testing.T, q *QuickList) {
	n := len(q.m)

	dis := map[[2]int]int{}

	src := q.head.next[0]

	for i := 0; i < n-1; i++ {
		dst := src
		for j := i + 1; j < n; j++ {
			dst = dst.next[0]
			dis[[2]int{src.key, dst.key}] = j - i
		}
		src = src.next[0]
	}

	q.Print()

	for levelIdx := 1; levelIdx < q.level; levelIdx++ {
		node := q.head.next[levelIdx]
		if node == nil {
			continue
		}
		for node.next[levelIdx] != nil {
			curKey := node.key
			curVal := node.val
			nextNode := node.next[levelIdx]

			nextKey := nextNode.key
			nextVal := nextNode.val
			if dis[[2]int{curKey, nextKey}] != node.span[levelIdx] {
				t.Fatalf("unexpected distance on level %d\nsrc kv:%v %v\ndst kv:%v %v\nexpect:%v\nget:%v\n",
					levelIdx, curKey, curVal, nextKey, nextVal, dis[[2]int{curKey, nextKey}], node.span[levelIdx])
			}
			node = node.next[levelIdx]
		}
	}
}

func TestConstructor(t *testing.T) {
	constructor := Constructor(10)
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())

	constructor.Leave(0)
	constructor.Leave(4)

	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())

	constructor.Leave(0)
	constructor.Leave(4)

	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())

	constructor.Leave(7)

	fmt.Println(constructor.Seat())

	constructor.Leave(3)

	fmt.Println(constructor.Seat())

	constructor.Leave(3)

	fmt.Println(constructor.Seat())

	constructor.Leave(9)

	fmt.Println(constructor.Seat())

	constructor.Leave(0)
	constructor.Leave(8)

	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())

	constructor.Leave(0)
	constructor.Leave(8)

	fmt.Println(constructor.Seat())
	fmt.Println(constructor.Seat())

	constructor.Leave(2)

}
