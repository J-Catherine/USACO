package main

import (
	"fmt"
)

var n,i,j,cntRight,cntLeft,maxCnt int
var s string


func main()  {
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	for i = 0; i < n; i++ {
		if s[(i-1+n)%n] != 'w' && s[i] != 'w' { //找断开点
			for j = i; j < i + n; j++ { //向右搜索直到不同
				if s[j % n] == s[i] || s[j%n] == 'w'{
					cntRight += 1
				} else{
					//fmt.Printf("right: %d %d  i %d",j%n,cntRight,i)
					break
				}
			}
			if cntRight == n {
				maxCnt = cntRight
				break
			}
			for j = i - 1; j >= i - n; j-- { //向左搜索直到不同
				if s[(j + n) %n] == s[(i-1+n)%n] || s[(j+n)%n] == 'w' {
					cntLeft += 1
				} else {
					//fmt.Printf(" left: %d %d\n",(j + n) %n,cntLeft)
					break
				}
			}
			if cntLeft + cntRight > maxCnt{
				maxCnt = cntLeft + cntRight
			}
			cntLeft = 0
			cntRight = 0
		}
	}
	fmt.Println(maxCnt)
}
