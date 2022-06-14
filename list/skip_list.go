package list

import (
	"fmt"
	"goal/constraint"
	"math/rand"
)

const MaxLevel = 32

type SkipList[K, V constraint.Number] struct {
	head, tail *SkipListNode[K, V]
	len        int
	level      int
	entries    [32]SkipListEntry[K, V]
}

func NewSkipList[K, V constraint.Number]() *SkipList[K, V] {
	res := &SkipList[K, V]{}
	res.len = 0
	res.head = newSkipListNode(*new(K), *new(V), 32)
	res.tail = res.head
	res.level = 1
	return res
}

func (l *SkipList[K, V]) Insert(key K, val V) {
	level := 1
	for rand.Intn(2) == 1 {
		level++
	}
	newNode := newSkipListNode(key, val, level)
	node := l.head
	maxLevel := max(32, max(l.level, level))
	levelIdx := maxLevel - 1
	pos := 0
	for levelIdx >= 0 {
		for node.level[levelIdx].next != nil && node.level[levelIdx].next.val <= val {
			pos += node.level[levelIdx].span
			node = node.level[levelIdx].next
		}
		l.entries[levelIdx].pos = pos
		l.entries[levelIdx].node = node
		levelIdx--
	}

	levelIdx = 0
	curNodePos := l.entries[0].pos + 1
	for levelIdx < level {
		node = l.entries[levelIdx].node
		nodePos := l.entries[levelIdx].pos
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
		if l.entries[levelIdx].node.level[levelIdx].next != nil {
			l.entries[levelIdx].node.level[levelIdx].span++
		}
		levelIdx++
	}

	newNode.back = l.entries[0].node

	l.level = maxLevel

}

func (l *SkipList[K, V]) Print() {

	levelIdx := 7
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

type SkipListNode[K, V constraint.Number] struct {
	key   K
	val   V
	level []SkipListLevel[K, V]
	back  *SkipListNode[K, V]
}

func newSkipListNode[K, V constraint.Number](key K, val V, level int) *SkipListNode[K, V] {
	res := &SkipListNode[K, V]{}
	res.key = key
	res.val = val
	res.level = make([]SkipListLevel[K, V], level)
	return res
}

type SkipListLevel[K, V constraint.Number] struct {
	next *SkipListNode[K, V]
	span int
}

type SkipListEntry[K, V constraint.Number] struct {
	node *SkipListNode[K, V]
	pos  int
}

func max[T constraint.Number](a, b T) T {
	if b > a {
		return b
	}
	return a
}
