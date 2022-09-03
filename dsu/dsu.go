package dsu

type DisjointSetUnion struct {
	fa []int
}

func NewDSU(n int) *DisjointSetUnion {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	return &DisjointSetUnion{arr}
}

func (d *DisjointSetUnion) Find(x int) int {
	if x != d.fa[x] {
		d.fa[x] = d.Find(d.fa[x])
	}
	return d.fa[x]
}

func (d *DisjointSetUnion) Union(x, y int) {
	fx := d.Find(x)
	fy := d.Find(y)
	if fx != fy {
		d.fa[fx] = fy
	}
}
