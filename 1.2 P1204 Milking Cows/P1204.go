package main

import (
	"fmt"
	"sort"
)

type Time struct {
	begin int
	end int
}
var (
	n,i, begin, end,ans1,ans2 int // ans1:最长有人  ans2：最长无人
	time            []Time
)

type timeSort []Time

func (a timeSort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a timeSort) Swap(i, j int){ // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a timeSort) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return a[i].begin < a[j].begin
}

func Max(a, b int) int  {
	if a > b { return a }
	return b
}

func Min(a,b int) int{
	if a < b { return a }
	return b
}

func main()  {
	fmt.Scanln(&n)
	time = make([]Time,n)
	for i =0; i< n;i++ {
		fmt.Scanln(&begin,&end)
		time[i].begin = begin
		time[i].end = end
	}
	sort.Sort(timeSort(time))  // go 结构体排序
	//fmt.Println(time)

	begin = time[0].begin
	end = time[0].end
	ans1 = end-begin  //初始化为第一个有人的时长
	ans2 = 0
	for i = 1; i < n; i++ {  //从第二个开始判断
		if time[i].begin <= end {
			if time[i].end >= end {
				end = time[i].end
				ans1 = Max(ans1, end - begin)
			}
		} else {
			ans1 = Max(ans1, time[i].end-time[i].begin)
			ans2 = Max(ans2, time[i].begin - end)
			begin = time[i].begin
			end = time[i].end
		}
	}
	fmt.Println(ans1, ans2)

}
