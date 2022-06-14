package list

import (
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	sl := NewSkipList[int, int]()
	sl.Insert(5, 1)
	sl.Insert(5, 2)
	sl.Insert(4, 1)
	sl.Insert(3, 1)
	sl.Insert(2, 1)
	sl.Insert(3, 2)
	sl.Insert(12, 1)
	sl.Insert(42, 1)
	sl.Insert(1, 1)
	sl.Insert(34, 1)
	sl.Insert(2, 2)

	sl.Print()
}
