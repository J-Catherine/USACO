package main

import (
	"fmt"
)

var (
	n, i, j, cntRight, cntLeft, maxCnt int
	before                             uint8
	s                                  string
)

func main() {
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	for i = 0; i < n; i++ {
		cntLeft = 0
		cntRight = 0
		if s[i] != 'w' { //找断开点，下一个不为w才可断开
			for j = i; j < i+n; j++ { //向右搜索直到不同
				if s[j%n] == s[i] || s[j%n] == 'w' {
					cntRight += 1
				} else {
					//fmt.Printf("right: %d %d  i %d\n", j%n, cntRight, i)
					break
				}
			}
			if cntRight == n {
				maxCnt = cntRight
				break
			}
			if s[(i-1+n)%n] != s[i] {
				if s[(i-1+n)%n] == 'w' { // 若上一个为w
					if s[i] == 'b' {
						before = 'r'
					} else {
						before = 'b'
					}
				} else {
					before = s[(i-1+n)%n]
				}
				for j = i-1; j >= i-n+cntRight; j-- { //向左搜索直到不同
					if s[(j+n)%n] == before || s[(j+n)%n] == 'w' {
						cntLeft += 1
					} else {
						//fmt.Printf(" left: %d %d\n", (j+n)%n, cntLeft)
						break
					}
				}
				if cntLeft+cntRight > maxCnt {
					maxCnt = cntLeft + cntRight
				}
			}
		}
	}
	if maxCnt == 0 {
		maxCnt = n
	}
	fmt.Println(maxCnt)
}
