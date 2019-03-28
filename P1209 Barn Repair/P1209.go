package main

import (
	"fmt"
	"sort"
)

var (
	m,s,c,i,j,l int
	num []int
	state [][]int
)

func main()  {
	fmt.Scanln(&m,&s,&c)
	num = make([]int, c)
	state = make([][]int,c,)
	for i=0; i<c; i++{
		num[i]=c
	}
	sort.Ints(num)
	for i,j = 0,1; j < c; i,j = i+1,j+1 {
		for num[i] == num[j] { j+=1 }
		num[i+1] = num[j]
	}
	l = i // 有牛的牛棚的个数
	for i=0
}