package list

import (
	"math/rand"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	//rand.Seed(time.Now().Unix())
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

	sl.head.key = -1
	sl.head.val = -1

	dis := map[[2]int]int{}

	src := sl.head

	for i := 0; i < n; i++ {
		dst := src
		for j := i + 1; j <= n; j++ {
			dst = dst.level[0].next
			//fmt.Println(src.val, dst.val)
			dis[[2]int{src.val, dst.val}] = j - i
		}
		src = src.level[0].next
	}

	sl.Print()

	for levelIdx := 1; levelIdx < sl.level; levelIdx++ {
		node := sl.head
		for node.level[levelIdx].next != nil {
			curVal := node.val
			nextVal := node.level[levelIdx].next.val
			if dis[[2]int{curVal, nextVal}] != node.level[levelIdx].span {
				t.Fatalf("not equal on level %d\nsrcv: %v, dstv: %v\nexpect:%v\nget:%v", levelIdx, curVal, nextVal, dis[[2]int{curVal, nextVal}], node.level[levelIdx].span)
			}
			node = node.level[levelIdx].next
		}
	}

}
