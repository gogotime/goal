package bit

// 生成n元集合所有k元子集的算法 https://leetcode.cn/problems/maximum-rows-covered-by-columns/?envType=daily-question&envId=Invalid%20Date
// Gosper's Hack算法 https://zhuanlan.zhihu.com/p/360512296
func maximumRows(matrix [][]int, numSelect int) int {
	mask := make([]int, len(matrix))
	for i := range matrix {
		for j := range matrix[i] {
			mask[i] |= matrix[i][j] << (len(matrix[i]) - 1 - j)
		}
	}

	cur := 1<<numSelect - 1
	limit := 1 << len(matrix[0])
	ans := 0

	for cur < limit {
		cnt := 0
		for i := range mask {
			if mask[i]|cur == cur {
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}

		lowbit := cur & (-cur)
		r := cur + lowbit
		cur = ((r^cur)>>2)/lowbit | r
	}

	return ans
}
