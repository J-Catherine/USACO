package main

import "fmt"

var (
	a,b,cnt,num []int
	n,res,non int
)

func main()  {
	fmt.Scanln(&n)
	a = make([]int,3)
	b = make([]int,3)
	num = make([]int,n)
	cnt = make([]int,3)
	if n<5{
		res = n*n*n*2
	} else{
		res = 250
	}
	for i:= 0; i<3; i++{
		fmt.Scan(&a[i])
	}
	for i:= 0; i<3; i++{
		fmt.Scan(&b[i])
	}
	//fmt.Println(a,b)
	non = 1
	for i:=0; i<3; i++{
		for j:= 0; j<5;j++{
			num[(b[i]-2+j+n)%n] = 1
		}
		for j:= 0; j<5;j++{
			if num[(a[i]-2+j+n)%n] == 1{
				cnt[i] += 1
				num[(a[i]-2+j+n)%n] = 0
			}
		}
		//fmt.Println(cnt[i])
		non *= cnt[i]
	}
	fmt.Println(res-non)
}
