package main

import (
	"fmt"
	"sort"
)

var (
	ans []int
	visited [][]int
	A,B,C,dep int
)

func min(x,y int) int {
	if x<y{
		return x
	}
	return y
}

func dfs(a,b int)  {
	dep++
	c := C-a-b
	//fmt.Println(dep,a,b,c)
	if visited[a][b] == 1{
		dep--
		return
	}
	visited[a][b] = 1

	//if a==1 && b==8 {for ;true;{;}}
	if a==0{
		ans = append(ans,c)
		//fmt.Printf("  %d %d %d",a,b,c)
	}
	if c > 0{
		if a<A{
			dfs(min(A,a+c),b)
		}
		if b<B{
			//fmt.Println("bB",b+c,b,c)
			dfs(a,min(B,b+c))
		}
	}
	if b >0 {
		if a<A { dfs(min(A,a+b), b-min(A,a+b)+a) }
		if c<C { dfs(a,b-min(C,c+b)+c) }
	}
	if a > 0{
		if b<B { dfs(a-min(B,a+b)+b, min(B,a+b)) }
		if c<C { dfs(a-min(C,a+c)+c, b) }
	}
	dep--
	return
}


func main()  {
	fmt.Scanln(&A,&B,&C)
	visited = make([][]int, 21)
	for i:=0; i<21; i++{
		visited[i] = make([]int,21)
	}
	ans = make([]int,0)
	dfs(0,0)
	sort.Ints(ans)
	for i, x :=range ans{
		if i == 0{
			fmt.Print(x)
		} else {
			fmt.Printf(" %d",x)
		}
	}
}
