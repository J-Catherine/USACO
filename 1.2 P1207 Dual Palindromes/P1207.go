package main

import (
	"fmt"
)

var (
	i,n,s,cnt,j,x,l,k,flag int
	t []int
)


func main()  {
	fmt.Scanln(&n,&s)  //前n个 >s的 且在两种或两种以上进制（二进制至十进制）上是回文数的十进制数
	t = make([]int,100)
	for i=s+1; n>0; i++{
		cnt = 0
		// 转2-10进制，并判断是否回文数
		for j=2; j<11; j++{ //j进制, 转2-10进制
			x = i
			flag=1
			for k := range(t){ t[k] = 0}

			for l=0; x>0; l++{ //l长度
				t[l] = x%j
				x/=j
			}
			for k=0;k<l/2;k++{  //判断是否回文数
				if t[k] != t[l-1-k]{
					flag = 0
					break
				}
			}
			if flag == 1 {
				cnt+=1
			}
			if cnt>1 {
				fmt.Println(i)
				n-=1
				break
			}
		}

	}
}
