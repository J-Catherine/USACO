package main

import (
	"fmt"
	"math"
)

var (
	//d1,d2,d3,d4,n,palindrome int
	start,end,n,cnt int
	m []int

)

func isPrime(x int) int {
	if x==2 || x==3 {
		return 1
	}
	if x== 1{
		return 0
	}
	if x%6!=1 && x%6!=5{
		return 0
	}
	n_sqrt := int(math.Sqrt(float64(x)))
	for i:=5;i<=n_sqrt;i+=6{
		if x%(i)==0 || x%(i+2)==0 {
				return 0
		}
	}
	return 1
}

func judge(x int)  {
	flag := 1
	for a:=x; a>0 && flag == 1; a/=10{
		flag = isPrime(a)
	}
	if flag == 1 {
		m[end] = x
		end+=1
	}
	return
}

//func fun(n int)  {
//	switch n {
//	case 1:
//		for d1 = 0; d1 <= 9; d1++ { //1
//			palindrome = d1;
//			judge(palindrome)
//		}
//	case 2:
//		for d1 = 0; d1 <= 9; d1++ { //2
//			palindrome = 10*d1 + d1
//			judge(palindrome)
//		}
//	case 3:
//		for d1 = 1; d1 <= 9; d1+=2 { // 3
//			for d2 = 0; d2 <= 9; d2++ {
//				palindrome = 100*d1 + 10*d2 + d1
//				judge(palindrome)
//			}
//		}
//	case 4:
//		for d1 = 1; d1 <= 9; d1+=2 { // 4
//			for d2 = 0; d2 <= 9; d2++ {
//				palindrome = 1000*d1 + 100*d2 + 10*d2 + d1
//				judge(palindrome)
//			}
//		}
//	case 5:
//		for d1 = 1; d1 <= 9; d1+=2 {    // 5
//			for d2 = 0; d2 <= 9; d2++ {
//				for d3 = 0; d3 <= 9; d3++ {
//					palindrome = 10000*d1 + 1000*d2 +100*d3 + 10*d2 + d1
//					judge(palindrome)
//				}
//			}
//		}
//	case 6:
//		for d1 = 1; d1 <= 9; d1+=2 {    // 6
//			for d2 = 0; d2 <= 9; d2++ {
//				for d3 = 0; d3 <= 9; d3++ {
//					palindrome = 100000*d1 + 10000*d2 +1000*d3 +100*d3 + 10*d2 + d1
//					judge(palindrome)
//				}
//			}
//		}
//	case 7:
//		for d1 = 1; d1 <= 9; d1+=2 {    // 7
//			for d2 = 0; d2 <= 9; d2++ {
//				for d3 = 0; d3 <= 9; d3++ {
//					for d4=0; d4<=9; d4++{
//						palindrome = 1000000*d1 + 100000*d2 +10000*d3 + 1000*d4 +100*d3 + 10*d2 + d1
//						judge(palindrome)
//					}
//
//				}
//			}
//		}
//	case 8:
//		for d1 = 1; d1 <= 9; d1+=2 {    // 8
//			for d2 = 0; d2 <= 9; d2++ {
//				for d3 = 0; d3 <= 9; d3++ {
//					for d4=0; d4<=9; d4++{
//						palindrome = 10000000*d1 + 1000000*d2 +100000*d3 + 10000*d4 + 1000*d4 +100*d3 + 10*d2 + d1
//						judge(palindrome)
//					}
//
//				}
//			}
//		}
//
//
//	}
//}

func main()  {
	fmt.Scanln(&n)
	//fun(n)
	m = make([]int, 1000000)
	v := 0
	end=1
	for i:=0; i<n; i++ {
		cnt = end-start
		for k:= start; k<start+cnt; k++{
			v= m[k]*10+1
			for j:=1; j<=9; j,v=j+1,v+1{
				judge(v)
			}
		}
		start += cnt

	}
	for i:=start; i<end; i++{
		fmt.Println(m[i])
	}

}
