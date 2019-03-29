package main

import (
	"fmt"
)

var (
	m,s,c,i,j,l,x int
	num,nums []int
	d [][]int
)

func min(x,y int) int {
	if x < y{
		return x
	}
	return y
}

func main()  {
	fmt.Scanln(&m,&s,&c)  //木板最大的数目M ,牛棚的总数S 和 牛的总数C(用空格分开)
	nums = make([]int, s+1)
	l = 0

	for i=1; i<c+1; i++{   // 读入有牛的牛棚的序号
		fmt.Scanln(&x)
		nums[x] = 1
	}
	for i=1; i<s+1; i++{
		if nums[i] ==1 { l+= 1}
	}

	num = make([]int, l+1)
	for i ,j= 1,1; i<=s; i++{
		if nums[i] == 1 {
			num[j] = i
			j += 1
		}
	}

	d = make([][]int,l+1)
	for i = 0; i< l+1; i++{
		d[i]=make([]int, m+1)
	}

	// 初始化
	for i = 1; i<l+1; i++{
		d[i][1] = num[i]-num[1]+1
	}


	for i =1; i<l+1; i++{
		for j=2; j<i && j<m+1; j++{
			d[i][j] = min(d[i-1][j-1]+1, d[i-1][j]+ num[i]-num[i-1])
		}
		if i <= m{
			d[i][i] = d[i-1][j-1]+1
		}
	}
	//for i =1; i<l+1; i++ {
	//	fmt.Println(d[i])
	//}
	fmt.Println(d[l][min(l,m)])
}