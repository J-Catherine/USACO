package main

import (
	"fmt"
	"io"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	list := make([]string, 201)
	now := ""
	_, _ = fmt.Scan(&now)
	listLength := 0
	for i := 0; now != "."; i, listLength = i+1, i+1 {
		list[i] = now
		_, _ = fmt.Scan(&now)
	}
	str := ""
	tmp := ""

	for {
		_, err := fmt.Scanln(&tmp)
		if err == io.EOF {
			break
		}
		str += tmp
	}

	//fmt.Println(str)

	f := make([]int, len(str))

	maxL := 0
	flag := 1
	for l := -1; l < len(str); l++ {
		if flag == 1 || f[l] == 1 {
			flag = 0
			left := l + 1
			for i := 0; i < listLength; i++ {
				if left+len(list[i])-1 < len(str) && list[i] == str[left:left+len(list[i])] {
					//fmt.Println(left + len(list[i]) - 1)
					f[left+len(list[i])-1] = 1
					maxL = max(left+len(list[i]), maxL)
				}
			}
		}
		if l == maxL {
			break
		}
	}
	fmt.Println(maxL)
}
