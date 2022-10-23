package st

import "math"

// ST表 《算法训练营进阶篇》p34
type ST struct {
	f [][]int
}

func NewST(arr []int) *ST {
	n := len(arr)
	k := int(math.Log2(float64(n)))
	f := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = make([]int, k+1)
		f[i][0] = arr[i-1]
	}

	for j := 1; j <= k; j++ {
		for i := 1; i <= n-(1<<j)+1; i++ {
			f[i][j] = max(f[i][j-1], f[i+(1<<(j-1))][j-1])
		}
	}
	return &ST{f: f}
}

func (st *ST) Query(l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	return max(st.f[l][k], st.f[r-(1<<k)+1][k])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ST表 树上倍增法求LCA问题 《算法训练营进阶篇》p45
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	d := map[*TreeNode]int{}
	id := []*TreeNode{}
	pa := map[*TreeNode]*TreeNode{}
	depth := 0
	qu := []*TreeNode{root}
	pa[root] = root
	for len(qu) > 0 {
		size := len(qu)
		for i := 0; i < size; i++ {
			v := qu[0]
			qu = qu[1:]
			d[v] = depth
			id = append(id, v)
			if v.Left != nil {
				qu = append(qu, v.Left)
				pa[v.Left] = v
			}
			if v.Right != nil {
				qu = append(qu, v.Right)
				pa[v.Right] = v
			}
		}
		depth++
	}

	k := int(math.Log2(float64(depth)))

	f := map[*TreeNode][]*TreeNode{}
	for c, par := range pa {
		f[c] = make([]*TreeNode, k+1)
		f[c][0] = par
	}

	for j := 1; j <= k; j++ {
		for _, c := range id {
			f[c][j] = f[f[c][j-1]][j-1]
		}
	}

	x, y := p, q
	if d[x] > d[y] {
		x, y = y, x
	}

	for j := k; j >= 0; j-- {
		v := f[y][j]
		if d[v] >= d[x] {
			y = v
		}
	}

	if x == y {
		return x
	}

	for j := k; j >= 0; j-- {
		v, u := f[x][j], f[y][j]
		if v != u {
			x, y = v, u
		}
	}
	return f[x][0]
}
