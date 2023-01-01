package list

import (
	"container/heap"
	"fmt"
	"math/rand"
)

const QuickListMaxLevel = 32
const pFactor = 0.25

type QuickList struct {
	level int
	m     map[int]int
	head  *QuickListNode
	prev  [QuickListMaxLevel]QuickListEntry
}

func NewQuickList() *QuickList {
	res := &QuickList{}
	res.head = newQuickListNode(-1, -1, 32)
	res.level = 1
	res.m = map[int]int{}
	return res
}

func (q *QuickList) Size() int {
	return len(q.m)
}

func (q *QuickList) FirstKey() int {
	return q.head.next[0].key
}

func (q *QuickList) LastKey() int {
	levelIdx := QuickListMaxLevel - 1
	node := q.head
	for levelIdx >= 0 {
		nextNode := node.next[levelIdx]
		for nextNode != nil {
			node = nextNode
			nextNode = node.next[levelIdx]
		}
		levelIdx--
	}
	return node.key
}

func (q *QuickList) Contains(key int) bool {
	_, ok := q.m[key]
	return ok
}

func (q *QuickList) NextKey(key int) int {
	val, ok := q.m[key]
	if !ok {
		return -1
	}
	q.searchPrev(key, val)
	next := q.prev[0].node.next[0].next[0]
	if next == nil {
		return -1
	}
	return next.key
}

func (q *QuickList) PrevKey(key int) int {
	val, ok := q.m[key]
	if !ok {
		return -1
	}
	q.searchPrev(key, val)
	return q.prev[0].node.key
}

func (q *QuickList) searchPrev(key, val int) int {
	levelIdx := QuickListMaxLevel - 1
	node := q.head
	pos := 0
	for levelIdx >= 0 {
		nextNode := node.next[levelIdx]
		for nextNode != nil && (nextNode.val < val || (nextNode.val == val && nextNode.key < key)) {
			pos += node.span[levelIdx]
			node = nextNode
			nextNode = node.next[levelIdx]
		}
		q.prev[levelIdx].pos = pos
		q.prev[levelIdx].node = node
		levelIdx--
	}
	return pos
}

func (q *QuickList) randomLevel() int {
	level := 1
	for level < QuickListMaxLevel && rand.Float64() < pFactor {
		level++
	}
	return level
}

func (q *QuickList) Insert(key, val int) {
	if _, ok := q.m[key]; ok {
		q.Delete(key)
	}
	level := q.randomLevel()
	newNode := newQuickListNode(key, val, level)

	maxLevel := max(q.level, level)
	levelIdx := maxLevel - 1

	q.searchPrev(key, val)

	levelIdx = 0
	curNodePos := q.prev[0].pos + 1

	for levelIdx < level {
		prev := q.prev[levelIdx].node
		prevPos := q.prev[levelIdx].pos
		next := prev.next[levelIdx]
		prev.next[levelIdx] = newNode
		newNode.next[levelIdx] = next
		if next != nil {
			if levelIdx == 0 {
				next.back = newNode
			}
			newNode.span[levelIdx] = prev.span[levelIdx] - (curNodePos - prevPos) + 1
		} else {
			newNode.span[levelIdx] = 0
		}
		prev.span[levelIdx] = curNodePos - prevPos
		levelIdx++
	}

	for levelIdx < maxLevel {
		if q.prev[levelIdx].node.next[levelIdx] != nil {
			q.prev[levelIdx].node.span[levelIdx]++
		}
		levelIdx++
	}

	newNode.back = q.prev[0].node

	q.level = maxLevel
	q.m[key] = val
}

func (q *QuickList) Add(key, val int) {
	if v, ok := q.m[key]; ok {
		q.Insert(key, v+val)
	} else {
		q.Insert(key, val)
	}
}

func (q *QuickList) Delete(key int) {
	val, ok := q.m[key]
	if !ok {
		return
	}

	q.searchPrev(key, val)

	levelIdx := 0
	cur := q.prev[0].node.next[0]
	for levelIdx < cur.level {
		prev := q.prev[levelIdx].node
		next := cur.next[levelIdx]

		prevSpan := prev.span[levelIdx]
		curSpan := cur.span[levelIdx]

		prev.next[levelIdx] = next
		prev.span[levelIdx] = prevSpan + curSpan - 1

		levelIdx++
	}

	for levelIdx < q.level {
		if q.prev[levelIdx].node.next[levelIdx] != nil {
			q.prev[levelIdx].node.span[levelIdx]--
		}
		levelIdx++
	}

	delete(q.m, key)
}

func (q *QuickList) Rank(key int) int {
	val, ok := q.m[key]
	if !ok {
		return -1
	}
	q.searchPrev(key, val)
	return q.prev[0].pos + q.prev[0].node.span[0] - 1
}

func (q *QuickList) Value(key int) (int, bool) {
	v, ok := q.m[key]
	return v, ok
}

func (q *QuickList) Print() {

	levelIdx := q.level - 1
	for levelIdx >= 0 {

		fmt.Printf("level %v:", levelIdx+1)
		fmt.Printf("head-- %v --", q.head.span[levelIdx])

		node := q.head.next[levelIdx]
		for node != nil {
			fmt.Printf("(%v,%v)", node.key, node.val)
			if node.span[levelIdx] != 0 {
				fmt.Printf("-- %v --", node.span[levelIdx])
			}
			node = node.next[levelIdx]
		}
		fmt.Println()
		levelIdx--
	}
	fmt.Println()

}

type QuickListNode struct {
	key, val int
	next     []*QuickListNode
	span     []int
	level    int
	back     *QuickListNode
}

func newQuickListNode(key, val int, level int) *QuickListNode {
	node := &QuickListNode{}
	node.key = key
	node.val = val
	node.level = level
	node.next = make([]*QuickListNode, level)
	node.span = make([]int, level)
	return node
}

type QuickListEntry struct {
	node *QuickListNode
	pos  int
}

type MaxDisHeap [][2]int

func (m *MaxDisHeap) Len() int {
	return len(*m)
}

func (m *MaxDisHeap) Less(i, j int) bool {
	i1, i2 := (*m)[i][0], (*m)[i][1]
	j1, j2 := (*m)[j][0], (*m)[j][1]
	d1 := (i2 - i1) / 2
	d2 := (j2 - j1) / 2
	return (d1 > d2) || (d1 == d2 && i1 < j1)
}

func (m *MaxDisHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxDisHeap) Push(x interface{}) {
	*m = append(*m, x.([2]int))
}

func (m *MaxDisHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return res
}

type ExamRoom struct {
	n int
	h *MaxDisHeap
	l *QuickList
}

func Constructor(n int) ExamRoom {
	h := &MaxDisHeap{}
	heap.Init(h)
	return ExamRoom{
		n: n,
		h: h,
		l: NewQuickList(),
	}
}

func (this *ExamRoom) Seat() int {
	if this.l.Size() == 0 {
		this.l.Add(0, 0)
		return 0
	}
	left, right := this.l.FirstKey(), this.n-1-this.l.LastKey()
	for this.l.Size() >= 2 {
		v := (*this.h)[0]
		p1, p2 := v[0], v[1]
		if this.l.Contains(p1) && this.l.Contains(p2) && this.l.NextKey(p1) == p2 {
			d := p2 - p1
			if d/2 <= left || d < right/2 {
				break
			}
			heap.Pop(this.h)
			this.l.Add(p1+d/2, p1+d/2)
			heap.Push(this.h, [2]int{p1, p1 + d/2})
			heap.Push(this.h, [2]int{p1 + d/2, p2})
			return p1 + d/2
		}
		heap.Pop(this.h)
	}
	if right > left {
		heap.Push(this.h, [2]int{this.l.LastKey(), this.n - 1})
		this.l.Add(this.n-1, this.n-1)
		return this.n - 1
	}
	this.l.Add(0, 0)
	heap.Push(this.h, [2]int{0, left})
	return 0
}

func (this *ExamRoom) Leave(p int) {
	if p != this.l.FirstKey() && p != this.l.LastKey() {
		prev := this.l.PrevKey(p)
		next := this.l.NextKey(p)
		heap.Push(this.h, [2]int{prev, next})
	}

	this.l.Delete(p)
}

//func max(a, b int) int {
//	if b < a {
//		return b
//	}
//	return a
//}
