package str

import "fmt"

// 最长回文子串 https://blog.csdn.net/weixin_43501684/article/details/124325764
func LongestPalindromicSubstring(s string) int {
	b := []byte{'^'}
	for i := range s {
		b = append(b, '#', s[i])
	}
	b = append(b, '#', '$')

	l := make([]int, len(b))
	l[1] = 1
	ans := 0
	k := 0
	r := 0
	for i := 2; i < len(b)-1; i++ {
		if i < r {
			l[i] = min(l[2*k-i], r-i)
		} else {
			l[i] = 1
		}

		for b[i-l[i]] == b[i+l[i]] {
			l[i]++
		}

		if i+l[i] > r {
			k = i
			r = i + l[i]
		}

		ans = max(ans, l[i]-1)
	}

	fmt.Println(string(b))
	fmt.Println(l)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
