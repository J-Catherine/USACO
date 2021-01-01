package main

import (
	"fmt"
	"math"
)

// 100,000,000
var (
	a,b,d1,d2,d3,d4,palindrome int
)

func judge(x int)  {
	if x<a ||x>b{
		return
	}

	if x==2 || x==3 {
		fmt.Println(x)
		return
	}
	if x%6!=1 && x%6!=5{
		return
	}
	n_sqrt := int(math.Sqrt(float64(x)))
	for i:=5;i<=n_sqrt;i+=6{
		if x%(i)==0 || x%(i+2)==0 {
			return
		}
	}

	fmt.Println(x)
	return
}

func main()  {
	fmt.Scanln(&a,&b)
	for d1 = 0; d1 <= 9; d1++ { //1
		palindrome = d1;
		judge(palindrome)
	}
	for d1 = 0; d1 <= 9; d1++ { //2
		palindrome = 10*d1 + d1
		judge(palindrome)
	}
	for d1 = 1; d1 <= 9; d1+=2 { // 3
		for d2 = 0; d2 <= 9; d2++ {
			palindrome = 100*d1 + 10*d2 + d1
			judge(palindrome)
		}
	}
	for d1 = 1; d1 <= 9; d1+=2 { // 4
		for d2 = 0; d2 <= 9; d2++ {
			palindrome = 1000*d1 + 100*d2 + 10*d2 + d1
			judge(palindrome)
		}
	}

	for d1 = 1; d1 <= 9; d1+=2 {    // 5
		for d2 = 0; d2 <= 9; d2++ {
			for d3 = 0; d3 <= 9; d3++ {
				palindrome = 10000*d1 + 1000*d2 +100*d3 + 10*d2 + d1
				judge(palindrome)
			}
		}
	}
	for d1 = 1; d1 <= 9; d1+=2 {    // 6
		for d2 = 0; d2 <= 9; d2++ {
			for d3 = 0; d3 <= 9; d3++ {
				palindrome = 100000*d1 + 10000*d2 +1000*d3 +100*d3 + 10*d2 + d1
				judge(palindrome)
			}
		}
	}
	for d1 = 1; d1 <= 9; d1+=2 {    // 7
		for d2 = 0; d2 <= 9; d2++ {
			for d3 = 0; d3 <= 9; d3++ {
				for d4=0; d4<=9; d4++{
					palindrome = 1000000*d1 + 100000*d2 +10000*d3 + 1000*d4 +100*d3 + 10*d2 + d1
					judge(palindrome)
				}

			}
		}
	}
	for d1 = 1; d1 <= 9; d1+=2 {    // 8
		for d2 = 0; d2 <= 9; d2++ {
			for d3 = 0; d3 <= 9; d3++ {
				for d4=0; d4<=9; d4++{
					palindrome = 10000000*d1 + 1000000*d2 +100000*d3 + 10000*d4 + 1000*d4 +100*d3 + 10*d2 + d1
					judge(palindrome)
				}

			}
		}
	}



}