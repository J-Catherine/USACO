package main

import "fmt"

var (
	n,i,j int
	a, b [][]uint8
	k uint8
)

func x90(a,b [][]uint8) int {
	for i = 0; i<n; i++{
		for j=0; j<n; j++{
			if b[i][j] != a[n-j-1][i]{
				return 0
			}
		}
	}
	return 1
}

func x180(a,b [][]uint8) int {
	for i = 0; i<n; i++{
		for j=0; j<n; j++{
			if b[i][j] != a[n-i-1][n-j-1]{
				return 0
			}
		}
	}
	return 1
}

func x270(a,b [][]uint8) int {
	for i = 0; i<n; i++{
		for j=0; j<n; j++{
			if b[i][j] != a[j][n-i-1]{
				return 0
			}
		}
	}
	return 1
}

func x0(a,b [][]uint8) int {
	for i = 0; i<n; i++{
		for j=0; j<n; j++{
			if b[i][j] != a[i][j]{
				return 0
			}
		}
	}
	return 1
}

func rotate(a [][]uint8) {
	for i = 0; i<n; i++{
		for j=0; j<n/2; j++{
			k = a[i][j]
			a[i][j] = a[i][n-j-1]
			a[i][n-j-1] = k
		}
	}
}

func main()  {
	fmt.Scanln(&n)
	a = make([][]uint8, n)
	b = make([][]uint8, n)
	for i=0; i<n; i++ { fmt.Scanln(&a[i])}
	for i=0; i<n; i++ { fmt.Scanln(&b[i])}
	if x90(a,b) == 1 {
		fmt.Println(1)
	} else {
		if x180(a,b) == 1 {
			fmt.Println(2)
		} else {
			if x270(a,b) == 1 {
				fmt.Println(3)
			} else {
				rotate(a)
				if x0(a,b) == 1{
					fmt.Println(4)
				} else {
					if x90(a,b) == 1 || x180(a,b) == 1 || x270(a,b) == 1 {
						fmt.Println(5)
					} else {
						rotate(a)
						if x0(a,b) == 1 {
							fmt.Println(6	)
						}else{
							fmt.Println(7)
						}
					}
				}
			}
		}
	}
}