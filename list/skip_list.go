package list

import (
	"fmt"
	"goal/constraint"
	"math/rand"
)

const MaxLevel = 32

type SkipList[K, V constraint.Comparable] struct {
	head, tail *SkipListNode[K, V]
	len        int
	level      int
	prevEntry  [32]SkipListEntry[K, V]
	m          map[K]V
}

func NewSkipList[K, V constraint.Comparable]() *SkipList[K, V] {
	res := &SkipList[K, V]{}
	res.len = 0
	res.head = newSkipListNode(*new(K), *new(V), 32)
	res.tail = res.head
	res.level = 1
	res.m = make(map[K]V)
	return res
}

func (l *SkipList[K, V]) searchPrev(key K, val V) int {
	levelIdx := MaxLevel - 1
	node := l.head
	pos := 0
	for levelIdx >= 0 {
		nextNode := node.level[levelIdx].next
		for nextNode != nil && (nextNode.val < val || (nextNode.val == val && nextNode.key < key)) {
			pos += node.level[levelIdx].span
			node = nextNode
			nextNode = node.level[levelIdx].next
		}
		l.prevEntry[levelIdx].pos = pos
		l.prevEntry[levelIdx].node = node
		levelIdx--
	}
	return pos
}

func (l *SkipList[K, V]) Insert(key K, val V) {
	if _, ok := l.m[key]; ok {
		l.Delete(key)
	}
	level := 1
	for rand.Intn(4) == 1 {
		level++
	}
	if level > MaxLevel {
		level = MaxLevel
	}
	newNode := newSkipListNode(key, val, level)

	maxLevel := max(l.level, level)
	levelIdx := maxLevel - 1
	l.searchPrev(key, val)

	levelIdx = 0
	curNodePos := l.prevEntry[0].pos + 1
	for levelIdx < level {
		node := l.prevEntry[levelIdx].node
		nodePos := l.prevEntry[levelIdx].pos
		next := node.level[levelIdx].next
		node.level[levelIdx].next = newNode
		newNode.level[levelIdx].next = next
		if next != nil {
			newNode.level[levelIdx].span = node.level[levelIdx].span - (curNodePos - nodePos) + 1
		} else {
			newNode.level[levelIdx].span = 0
		}
		node.level[levelIdx].span = curNodePos - nodePos
		levelIdx++
	}

	for levelIdx < maxLevel {
		if l.prevEntry[levelIdx].node.level[levelIdx].next != nil {
			l.prevEntry[levelIdx].node.level[levelIdx].span++
		}
		levelIdx++
	}

	newNode.back = l.prevEntry[0].node

	l.level = maxLevel
	l.m[key] = val
}

func (l *SkipList[K, V]) Add(key K, val V) {
	if v, ok := l.m[key]; ok {
		l.Insert(key, v+val)
	} else {
		l.Insert(key, val)
	}
}

func (l *SkipList[K, V]) Delete(key K) {
	val, ok := l.m[key]
	if !ok {
		return
	}

	l.searchPrev(key, val)

	levelIdx := 0
	cur := l.prevEntry[0].node.level[0].next
	for levelIdx < len(cur.level) {
		prev := l.prevEntry[levelIdx].node
		next := cur.level[levelIdx].next

		prevSpan := prev.level[levelIdx].span
		curSpan := cur.level[levelIdx].span

		prev.level[levelIdx].next = next
		prev.level[levelIdx].span = prevSpan + curSpan - 1

		levelIdx++
	}

	for levelIdx < l.level {
		if l.prevEntry[levelIdx].node.level[levelIdx].next != nil {
			l.prevEntry[levelIdx].node.level[levelIdx].span--
		}
		levelIdx++
	}

	delete(l.m, key)
}

func (l *SkipList[K, V]) Value(key K) (V, bool) {
	v, ok := l.m[key]
	return v, ok
}

func (l *SkipList[K, V]) Rank(key K) int {
	val, ok := l.m[key]
	if !ok {
		return -1
	}
	l.searchPrev(key, val)
	return l.prevEntry[0].pos
}

func (l *SkipList[K, V]) Print() {

	levelIdx := l.level - 1
	for levelIdx >= 0 {

		fmt.Printf("level %v:", levelIdx+1)
		fmt.Printf("head--%v--", l.head.level[levelIdx].span)

		node := l.head.level[levelIdx].next
		for node != nil {
			fmt.Printf("(%v,%v)", node.key, node.val)
			if node.level[levelIdx].span != 0 {
				fmt.Printf("- %v -", node.level[levelIdx].span)
			}
			node = node.level[levelIdx].next
		}
		fmt.Println()
		levelIdx--
	}
	fmt.Println()

}

type SkipListNode[K, V constraint.Comparable] struct {
	key   K
	val   V
	level []SkipListLevel[K, V]
	back  *SkipListNode[K, V]
}

func newSkipListNode[K, V constraint.Comparable](key K, val V, level int) *SkipListNode[K, V] {
	res := &SkipListNode[K, V]{}
	res.key = key
	res.val = val
	res.level = make([]SkipListLevel[K, V], level)
	return res
}

type SkipListLevel[K, V constraint.Comparable] struct {
	next *SkipListNode[K, V]
	span int
}

type SkipListEntry[K, V constraint.Comparable] struct {
	node *SkipListNode[K, V]
	pos  int
}

func max[T constraint.Comparable](a, b T) T {
	if b > a {
		return b
	}
	return a
}

func min[T constraint.Comparable](a, b T) T {
	if b < a {
		return b
	}
	return a
}
