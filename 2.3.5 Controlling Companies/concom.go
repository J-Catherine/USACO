package main

import "fmt"

var (
	n,num int
	visited,stock,own []int
	com [][]int
)

func max(i,j int) int {
	if i>j{
		return i
	}
	return j
}

func dfs(i int)  {
	if visited[i] == 1{
		return
	}
	visited[i] = 1

	for j:=1; j<= num; j++ {  //枚举所有公司
		stock[j] += com[i][j]
		if stock[j] > 50{
			own[j] = 1
			dfs(j)
		}
	}

}

func main()  {
	_, _ = fmt.Scanln(&n)
	com = make([][]int, 201)
	for i:=0; i<201; i++{
		com[i] = make([]int, 201)
	}
	for k:=1; k<=n; k++{
		i,j,p := 0,0,0
		_,_ = fmt.Scanln(&i, &j, &p)

		com[i][j] = p
		num = max(i, max(num,j))  //公司个数
	}
	//fmt.Println(com)


	for i:=1; i<=num; i++{
		own = make([]int, num+1)  // 当前阶段i公司能否控制j公司
		visited = make([]int, num+1)   // 当前阶段j公司有无被访问过
		stock = make([]int, num+1)    // 当前阶段i公司累积用用j公司的股份
		dfs(i)
		for j:=1; j<=num; j++{
			if own[j] == 1 && i!=j {
				fmt.Println(i,j)
			}
		}
	}

}
