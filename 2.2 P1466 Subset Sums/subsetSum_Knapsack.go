package main

import "fmt"

func main()  {
	// F(i,j) = F(i-1, j) + F(i,j-i)
	n:=0
	fmt.Scanln(&n)
	f := make([][]int,n+1)
	sum := (1+n)*n/2
	if sum %2 != 0{
		fmt.Println(0)
		return
	}
	sum /= 2

	for i:=0; i<=n; i++{
		f[i] = make([]int, sum+1)
	}

	f[0][0] = 1
	for i:=1; i<=n; i++{
		for j:=0; j<=sum; j++{
			if j>=i {
				f[i][j] = f[i-1][j] + f[i-1][j-i]
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}
	fmt.Println(f[n][sum]/2)
}