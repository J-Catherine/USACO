package main

import (
	"fmt"
)

var (
	m, s, c, i, j, l, x int
	num, nums           []int
	d                   [][]int
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Scanln(&m, &s, &c) //木板最大的数目M ,牛棚的总数S 和 牛的总数C(用空格分开)
	nums = make([]int, s+1)
	l = 0 // 有牛的牛棚数

	for i = 1; i < c+1; i++ { // 读入有牛的牛棚的序号，注意多头牛可能在一个牛棚中
		fmt.Scanln(&x)
		nums[x] = 1
	}
	for i = 1; i < s+1; i++ {
		if nums[i] == 1 {
			l += 1
		}
	}
	num = make([]int, l+1)
	for i, j = 1, 1; i <= s; i++ {
		if nums[i] == 1 {
			num[j] = i
			j += 1
		}
	}

	// 动归，d[i][j] 前i个分成j段 所需要木材的最小值
	// d[i][j] = min{d[i-1][j-1]+1, d[i-1][j]+num[i]-num[i-1]}, i>m
	//         = d[i-1][j-1] + 1 , i<=m 有足够的木板
	d = make([][]int, l+1)
	for i = 0; i < l+1; i++ {
		d[i] = make([]int, m+1)
	}

	// 初始化
	for i = 1; i < l+1; i++ {
		d[i][1] = num[i] - num[1] + 1
	}

	for i = 1; i < l+1; i++ {
		for j = 2; j < i && j < m+1; j++ {
			d[i][j] = min(d[i-1][j-1]+1, d[i-1][j]+num[i]-num[i-1])
		}
		if i <= m {
			d[i][i] = d[i-1][j-1] + 1
		}
	}
	fmt.Println(d[l][min(l, m)])
}
