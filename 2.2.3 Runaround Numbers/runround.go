package main

import (
	"fmt"
	"strconv"
)

func isRunaround(n int) int {
	snum := strconv.Itoa(n)
	length := len(snum)
	move := make([]int, length)
	flag := make([]int, 10)
	for i := 0; i < length; i++ {
		move[i] = int(snum[i]) - 48
	}
	p := 0
	flag[0] = 1
	for i := 0; i < length; i++ {
		if flag[move[p]] == 1 {
			return 0
		}
		flag[move[p]] = 1
		p = (p + move[p]) % length
	}
	if p != 0 {
		return 0
	}
	fmt.Println(n)
	return 1
}

func main() {
	n := 0
	fmt.Scanln(&n)
	for n = n + 1; isRunaround(n) != 1; n++ {
	}
}
