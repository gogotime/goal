package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTreeNode_PreOrderTraversal(t *testing.T) {
	root := newBinaryTreeFromArray([]int{5, -1, 2, 4, -1, 3, 6}, -1)
	root.MorisPre(func(v int) {
		fmt.Println(v)
	})
	fmt.Println("-------------")
	root.MorisIn(func(v int) {
		fmt.Println(v)
	})
	//root = nil
	//root.PreOrderTraversal(nil)
}

func TestBinaryTreeNode_DeleteNode(t *testing.T) {
	root := newBinaryTreeFromArray([]int{5, -1, 2, 4, -1, 3, 6}, -1)
	root.MorisIn(func(v int) {
		fmt.Println(v)
	})
	fmt.Println("--------------")
	root.DeleteNode(root.Value, func(v1, v2 int) bool {
		return v1 < v2
	}).MorisIn(func(v int) { fmt.Println(v) })

}
