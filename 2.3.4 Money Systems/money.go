package main

import "fmt"

func main() {
	v, n := 0,0
	_, _ = fmt.Scanln(&v, &n)
	h := make([]int,v)
	for i := 0; i < v; i++ {
		fmt.Scan(&h[i])
	}

	f:= make([]int,n+1)
	f[0]=1
	for i:=0; i<v; i++{
		for j:= h[i]; j<=n; j++{
			f[j] += f[j-h[i]]
			//fmt.Println(j,j-h[i], f[j])
		}
	}
	fmt.Println(f[n])
}
