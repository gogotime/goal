package tree

import (
	"fmt"
	"goal/constraint"
)

type BinaryTreeNode[T any] struct {
	Left, Right *BinaryTreeNode[T]
	Value       T
}

func newBinaryTree[T any](value T) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{Value: value}
}

func newBinaryTreeFromArray[T constraint.Comparable](arr []T, nilValue T) *BinaryTreeNode[T] {
	if len(arr) == 0 || arr[0] == nilValue {
		return newBinaryTree[T](*new(T))
	}
	root := &BinaryTreeNode[T]{Value: arr[0]}
	q := []*BinaryTreeNode[T]{root}
	i := 1
	for len(q) != 0 {
		size := len(q)
		for j := 0; j < size; j++ {
			v := q[0]
			q = q[1:]
			fmt.Println(v, i)
			if i >= len(arr) {
				break
			}
			if arr[i] != nilValue {
				v.Left = &BinaryTreeNode[T]{Value: arr[i]}
				q = append(q, v.Left)
			}
			i++
			if i >= len(arr) {
				break
			}
			if arr[i] != nilValue {
				v.Right = &BinaryTreeNode[T]{Value: arr[i]}
				q = append(q, v.Right)
			}
			i++
		}
	}
	return root
	// TODO morris
}

type RBTreeNode[T constraint.Number] struct {
	BinaryTreeNode[T]
	R bool
}

func (b *BinaryTreeNode[T]) MorisPre(handler func(v T)) {
	cur := b
	for cur != nil {
		if cur.Left != nil {
			morisRight := cur.Left
			for morisRight.Right != nil && morisRight.Right != cur {
				morisRight = morisRight.Right
			}
			if morisRight.Right == nil {
				morisRight.Right = cur
				handler(cur.Value)
				cur = cur.Left
				continue
			} else {
				morisRight.Right = nil
			}
		} else {
			handler(cur.Value)
		}
		cur = cur.Right
	}
}

func (b *BinaryTreeNode[T]) MorisIn(handler func(v T)) {
	cur := b
	for cur != nil {
		if cur.Left != nil {
			morisRight := cur.Left
			for morisRight.Right != nil && morisRight.Right != cur {
				morisRight = morisRight.Right
			}
			if morisRight.Right == nil {
				morisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				handler(cur.Value)
				morisRight.Right = nil
			}
		} else {
			handler(cur.Value)
		}
		cur = cur.Right
	}
}

func (b *BinaryTreeNode[T]) DeleteNode(key T, less func(v1, v2 T) bool) *BinaryTreeNode[T] {
	switch {
	case less(b.Value, key):
		b.Right = b.DeleteNode(key, less)
	case less(key, b.Value):
		b.Left = b.DeleteNode(key, less)
	case b.Left == nil || b.Right == nil:
		if b.Left == nil {
			return b.Right
		}
		return b.Left
	default:
		resNode := b.Right
		for resNode.Left != nil {
			resNode = resNode.Left
		}
		resNode.Right = b.Right.DeleteNode(resNode.Value, less)
		resNode.Left = b.Left
		return resNode
	}
	return b
}
