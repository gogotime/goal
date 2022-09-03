package priorityqueue

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	type Wrapper struct {
		i int
	}
	queue := NewPriorityQueue(func(a, b *Wrapper) bool {
		return a.i <= b.i
	})
	fmt.Println(queue.Pop())
	queue.Push(&Wrapper{5})
	queue.Push(&Wrapper{3})
	queue.Push(&Wrapper{4})
	queue.Push(&Wrapper{1})
	queue.Push(&Wrapper{2})
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
