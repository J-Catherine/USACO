package main

import "fmt"

var (
	sum, ans int
	a        []int
)

func dfs(now, curSum int, tmp []int) {
	if now < 0 {
		return
	}
	if now == 0 && curSum == sum {
		ans += 1
		//fmt.Println(tmp)
		return
	}
	dfs(now-1, curSum, tmp)
	dfs(now-1, curSum+now, append(tmp, now))
}

func main() {
	n := 0
	_, _ = fmt.Scanln(&n)
	sum = (1 + n) * n / 2
	if sum%2 != 0 {
		fmt.Println(0)
	} else {
		sum = sum / 2
		a = make([]int, 0)
		dfs(n, 0, a)
		fmt.Println(ans / 2)
	}
}
