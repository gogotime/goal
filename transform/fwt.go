package transform

// 快速沃尔什变换
func fwt(arr []int, rev bool) {
	for mid := 1; mid < len(arr); mid <<= 1 {
		for block, j := mid<<1, 0; j < len(arr); j += block {
			for k := 0; k < mid; k++ {
				if !rev {
					arr[j+k] += arr[j+k+mid]
				} else {
					arr[j+k] -= arr[j+k+mid]
				}
			}
		}
	}
}

// 例题:按位与为零的三元组 https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/submissions/
func countTriplets(nums []int) int {
	arr := make([]int, 1<<16)
	for _, v := range nums {
		arr[v]++
	}

	fwt(arr, false)
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i] * arr[i]
	}
	fwt(arr, true)
	return arr[0]
}
