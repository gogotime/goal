package priorityqueue

type PriorityQueue[T any] struct {
	arr  []*T
	less Less[T]
}

type Less[T any] func(a, b *T) bool

func NewPriorityQueue[T any](less Less[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{less: less}
}

func (p *PriorityQueue[T]) Push(v *T) {
	if v == nil {
		return
	}
	p.arr = append(p.arr, v)
	i := len(p.arr)
	for i > 1 {
		fa := i / 2
		j, k := i-1, fa-1
		if p.less(p.arr[j], p.arr[k]) {
			p.arr[j], p.arr[k] = p.arr[k], p.arr[j]
			i = fa
			continue
		}
		break
	}
}

func (p *PriorityQueue[T]) Pop() *T {
	if len(p.arr) == 0 {
		return nil
	}
	res := p.arr[0]
	p.arr[0] = p.arr[len(p.arr)-1]
	p.arr = p.arr[:len(p.arr)-1]
	i := 1
	n := len(p.arr)
	for i <= n {
		li, ri := 2*i, 2*i+1
		j, k := i-1, li-1
		if li <= n && !p.less(p.arr[j], p.arr[k]) {
			p.arr[j], p.arr[k] = p.arr[k], p.arr[j]
			i = li
			continue
		}
		k = ri - 1
		if ri <= n && !p.less(p.arr[i], p.arr[ri]) {
			p.arr[i], p.arr[ri] = p.arr[ri], p.arr[i]
			i = ri
			continue
		}
		break
	}

	return res
}
