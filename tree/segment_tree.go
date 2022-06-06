package tree

type segTree[T any] struct {
	root       *segTreeNode[T]
	selectFunc func(lv, rv T) T
	updateFunc func(val, add T) T
	start, end int
}

func NewSegTree[T any](start, end int, selectFunc func(lv, rv T) T, updateFunc func(val, add T) T) *segTree[T] {
	return &segTree[T]{
		root:       newSegTreeNode[T](),
		selectFunc: selectFunc,
		updateFunc: updateFunc,
		start:      start,
		end:        end,
	}
}

func (t *segTree[T]) Update(l, r int, v T) {
	segUpdate(t.root, t.start, t.end, l, r, v, t.selectFunc, t.updateFunc)
}

func (t *segTree[T]) Query(l, r int) T {
	return segQuery(t.root, t.start, t.end, l, r, t.selectFunc, t.updateFunc)
}

func (t *segTree[T]) GetRootVal() T {
	return t.root.val
}

type segTreeNode[T any] struct {
	leftChild, rightChild *segTreeNode[T]
	val, add              T
	lazy                  bool
}

func newSegTreeNode[T any]() *segTreeNode[T] {
	return &segTreeNode[T]{}
}

func segUpdate[T any](node *segTreeNode[T], curL, curR, ragL, ragR int, v T, selectFunc func(lv, rv T) T, updateFunc func(val, add T) T) {
	if ragL <= curL && curR <= ragR {
		node.add = updateFunc(node.add, v)
		node.val = updateFunc(node.val, v)
		node.lazy = true
		return
	}
	segPushDown(node, updateFunc)
	mid := (curL + curR) / 2
	if ragL <= mid {
		segUpdate(node.leftChild, curL, mid, ragL, ragR, v, selectFunc, updateFunc)
	}
	if ragR > mid {
		segUpdate(node.rightChild, mid+1, curR, ragL, ragR, v, selectFunc, updateFunc)
	}
	segPushUp(node, selectFunc)
}

func segQuery[T any](node *segTreeNode[T], curL, curR, ragL, ragR int, selectFunc func(lv, rv T) T, updateFunc func(val, add T) T) T {
	if ragL <= curL && curR <= ragR {
		return node.val
	}
	segPushDown(node, updateFunc)
	mid := (curL + curR) / 2
	lv, rv := *new(T), *new(T)
	if ragL <= mid {
		lv = segQuery[T](node.leftChild, curL, mid, ragL, ragR, selectFunc, updateFunc)
	}
	if ragR > mid {
		rv = segQuery[T](node.rightChild, mid+1, curR, ragL, ragR, selectFunc, updateFunc)
	}
	return selectFunc(lv, rv)
}

func segPushDown[T any](node *segTreeNode[T], updateFunc func(val, add T) T) {
	if node.leftChild == nil {
		node.leftChild = newSegTreeNode[T]()
	}
	if node.rightChild == nil {
		node.rightChild = newSegTreeNode[T]()
	}
	if !node.lazy {
		return
	}
	node.leftChild.add = updateFunc(node.leftChild.add, node.add)
	node.leftChild.val = updateFunc(node.leftChild.val, node.add)
	node.leftChild.lazy = true
	node.rightChild.add = updateFunc(node.rightChild.add, node.add)
	node.rightChild.val = updateFunc(node.rightChild.val, node.add)
	node.rightChild.lazy = true
	node.add = *new(T)
	node.lazy = false
}

func segPushUp[T any](node *segTreeNode[T], selectFunc func(lv, rv T) T) {
	node.val = selectFunc(node.leftChild.val, node.rightChild.val)
}
