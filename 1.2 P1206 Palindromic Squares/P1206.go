package main

import (
	"fmt"
)

var(
	i,j,n,x int
	t,ti []int
	d []uint8
)

func isPalindromic(i int) {
	x = i*i
	l := 0
	li := 0
	for index := range(ti){
		ti[index] = 0
	}
	for j =0; i>0; j++{
		ti[j] = i%n
		i /= n
		li += 1
	}

	for index := range(t){
		t[index] = 0
	}
	for j=0; x>0; j++{
		t[j] = x%n
		x /= n
		l += 1
	}
	for j = 0; j<l/2; j++ {
		if t[j] != t[l-1-j] {
			return
		}
	}
	for j = li-1; j>-1; j--{
		fmt.Printf("%c",d[ti[j]])
	}
	fmt.Print(" ")
	for j = l-1; j >-1; j--{
		fmt.Printf("%c",d[t[j]])
	}
	fmt.Println()
}




func main()  {
	fmt.Scanln(&n)
	t = make([]int, 100)
	ti = make([]int, 16)
	d = []uint8 {'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','G','H','I','J','K'}
	for i = 0; i < 300; i++ {
		isPalindromic(i+1) //判断是否是回文数
	}
}
