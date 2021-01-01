package main

import (
	"fmt"
)

var (
	n int
	ans [][]int
)

func operation(way int) string {
	str :=""
	switch way {
		case 0:
			str = "+"
		case 1:
			str = "-"
		case 2:
			str = " "
	}
	return str
}

func dfs(cur, sum, last int, res []int)  {
	if cur > n{
		sum += last
		if sum == 0 {
			tmp := make([]int, n+1)
			copy(tmp, res)
			ans = append(ans, tmp)
			//fmt.Println(res)
		}
		return
	}
	res[cur] = 0
	dfs(cur+1, sum+last,cur, res)

	res[cur] = 1
	dfs(cur+1, sum+last, -cur, res)

	res[cur] = 2
	tmp := 0
	if last < 0 {
		tmp = -1 * ((-last)*10 + cur)
	} else {
		tmp = last*10 + cur
	}
	dfs(cur+1, sum, tmp, res)
}

func main()  {
	_,_ = fmt.Scanln(&n)
	res := make([]int,n+1)
	//ans = make([][]int, 81*81+1)
	dfs(2, 0, 1, res)
	//fmt.Println(ans)


	str := make([]string, len(ans))
	for i:=0; i<len(ans); i++{
		str[i] = "1"
		for j:=2; j< n+1; j++{
			//fmt.Println(operation(ans[i][j]))
			str[i] += operation(ans[i][j]) + string(j+48)
		}
	}
	for i:=0; i<len(ans); i++{
		for j:=i+1; j<len(ans); j++{
			if str[i] > str[j]{
				str[i], str[j] = str[j], str[i]
			}
		}
	}
	for i:=0; i<len(ans); i++ {
		fmt.Println(str[i])
	}
}
