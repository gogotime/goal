package geometry

import "sort"

func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

// Jarvis2D returns the convex hull of points
// refers: leetcode.587 https://leetcode.cn/problems/erect-the-fence/
func Jarvis2D(points [][]int) [][]int {
	n := len(points)
	if n <= 3 {
		return points
	}

	lm := 0
	for i, v := range points {
		if v[0] < points[lm][0] || (v[0] == points[lm][0] && v[1] < points[lm][1]) {
			lm = i
		}
	}

	vis := make([]bool, n)
	vis[lm] = true

	p := lm
	ans := [][]int{points[lm]}
	for {
		q := (p + 1) % n
		for i := range points {
			if cross(points[p], points[q], points[i]) < 0 {
				q = i
			}
		}

		for i := range points {
			if !vis[i] && i != p && i != q && cross(points[p], points[q], points[i]) == 0 {
				vis[i] = true
				ans = append(ans, points[i])
			}
		}

		if !vis[q] {
			vis[q] = true
			ans = append(ans, points[q])
		}
		p = q
		if p == lm {
			break
		}
	}
	return ans
}

func distance(p, q []int) int {
	return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
}

// Graham2D returns the convex hull of points
// refers: leetcode.587 https://leetcode.cn/problems/erect-the-fence/
func Graham2D(points [][]int) [][]int {
	n := len(points)
	if n <= 3 {
		return points
	}
	lm := 0
	for i := range points {
		if points[i][0] < points[lm][0] {
			lm = i
		}
	}

	points[0], points[lm] = points[lm], points[0]

	bottom := points[0]
	tr := points[1:]
	sort.Slice(tr, func(i, j int) bool {
		c := cross(bottom, tr[i], tr[j])
		return c > 0 || (c == 0 && distance(bottom, tr[i]) < distance(bottom, tr[j]))
	})

	i := n - 1
	for i >= 0 && cross(bottom, points[n-1], points[i]) == 0 {
		i--
	}

	l, r := i+1, n-1
	for l < r {
		points[l], points[r] = points[r], points[l]
		l++
		r--
	}

	stk := []int{0, 1}
	for i = 2; i < n; i++ {
		for len(stk) > 1 && cross(points[stk[len(stk)-2]], points[stk[len(stk)-1]], points[i]) < 0 {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}

	ans := [][]int{}
	for _, v := range stk {
		ans = append(ans, points[v])
	}
	return ans
}
