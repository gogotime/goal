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
