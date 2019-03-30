package main

import (
	"fmt"
	"sort"
)

var (
	p,q,n,m,a,b,flag,cnt,nums int
	num,pos []int
)

func main()  {
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	num =make([]int,125001)
	nums = 0
	for p=0;p<=m;p++{
		for q=0;q<=p;q++{
			if num[p*p+q*q] == 0{
				pos=append(pos,p*p+q*q)
				nums+=1
			}
			num[p*p+q*q] =1
		}
	}
	sort.Ints(pos)

	for b=1; (n-1)*b <= m*m*2; b++{
		for i:=0; i<nums && pos[i]+(n-1)*b <= m*m*2; i++{
			a = pos[i]
			flag = 1
			for j := 0; j < n; j++ {
				if num[a+j*b] != 1 {
					flag = 0
				}
				if flag == 0 {
					break
				}
			}
			if flag == 1 {
				cnt+=1
				fmt.Println(a, b)
			}
		}
	}
	if cnt == 0{
		fmt.Println("NONE")
	}
}
