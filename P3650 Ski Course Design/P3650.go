package main

import (
	"fmt"
	"sort"
)

var (
	n,sum,res int
	num []int
)

func min(x,y int) int {
	if x<y{
		return x
	}
	return y
}

func  main()  {
	fmt.Scanln(&n)
	num = make([]int, n)
	res = 1000000000
	for i:=0; i<n; i++{
		fmt.Scanln(&num[i])
	}
	sort.Ints(num)
	//fmt.Println(num)
	for h:=num[0]; h<=num[n-1]-17; h++{
		sum = 0
		for l:=0; l<n && h-num[l]>0;l++{
			sum += (h-num[l])*(h-num[l])
		}

		for r:=n-1; r>=0 && num[r]-h > 17; r--{
			sum+= (num[r]-h-17)*(num[r]-h-17)
		}
		res = min(res,sum)
	}
	fmt.Println(res)
}
