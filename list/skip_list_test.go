package list

import (
	"fmt"
	"goal/constraint"
	"math/rand"
	"testing"
)

func init() {
	//rand.Seed(time.Now().Unix())
}

func TestSkipList_Insert(t *testing.T) {
	sl := newRandSkipList(1000)

	checkDistance(t, sl)
}

func BenchmarkSkipList_Insert(b *testing.B) {
	sl := NewSkipList[int, int]()
	v := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		v[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Insert(v[i], v[i])
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := NewSkipList[int, int]()
	only := map[int]int{}
	n := 1000
	mod := 1000000
	for i := 0; i < n; {
		r := rand.Int() % mod
		if _, ok := only[r]; !ok {
			only[r] = r
			sl.Insert(r, r)
			i++
		}
	}

	i := 0
	for k := range sl.m {
		if i == n/2 {
			break
		}
		sl.Delete(k)
		i++
	}

	checkDistance(t, sl)
}

func BenchmarkSkipList_Delete(b *testing.B) {
	sl := newRandSkipList(b.N)
	for i := 0; i < b.N; i++ {
		sl.Delete(i)
	}
}

func TestSkipList_Rank(t *testing.T) {
	sl := newRandSkipList(100)
	sl.Print()
	fmt.Println(sl.Rank(43))
}

func newRandSkipList(n int) *SkipList[int, int] {
	sl := NewSkipList[int, int]()
	for i := 0; i < n; i++ {
		sl.Insert(i, rand.Int()%n)
	}
	return sl
}

func checkDistance[K, V constraint.Comparable](t *testing.T, sl *SkipList[K, V]) {
	n := len(sl.m)

	dis := map[[2]K]int{}

	src := sl.head.level[0].next

	for i := 0; i < n-1; i++ {
		dst := src
		for j := i + 1; j < n; j++ {
			dst = dst.level[0].next
			dis[[2]K{src.key, dst.key}] = j - i
		}
		src = src.level[0].next
	}

	sl.Print()

	for levelIdx := 1; levelIdx < sl.level; levelIdx++ {
		node := sl.head.level[levelIdx].next
		if node == nil {
			continue
		}
		for node.level[levelIdx].next != nil {
			curKey := node.key
			curVal := node.val
			nextNode := node.level[levelIdx].next

			nextKey := nextNode.key
			nextVal := nextNode.val
			if dis[[2]K{curKey, nextKey}] != node.level[levelIdx].span {
				t.Fatalf("unexpected distance on level %d\nsrc kv:%v %v\ndst kv:%v %v\nexpect:%v\nget:%v\n",
					levelIdx, curKey, curVal, nextKey, nextVal, dis[[2]K{curKey, nextKey}], node.level[levelIdx].span)
			}
			node = node.level[levelIdx].next
		}
	}
}
