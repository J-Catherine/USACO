package main

import (
	"fmt"
)

var (
	n, i,ans,c  int
	left, right[][]int
	s                                  string
)


func getLeft(i int, col uint8, from int)  int{
	if col == 'r'{
		c = 0
	} else {
		c =1
	}
	if left[c][i] == -1 {
		if col != s[i] && s[i]!='w'{
			left[c][i] = 0
			return 0
		}else {
			if from != i {
				left[c][i] = getLeft((i-1+n)%n, col, from) + 1
			}else{
				left[c][i] = n
			}
		}
	}
	return left[c][i]
}

func getRight(i int, col uint8, from int)  int{
	if col == 'r'{
		c = 0
	} else {
		c = 1
	}
	if right[c][i] == -1 {
		if col != s[i] && s[i]!='w'{
			right[c][i] = 0
			return 0
		}else {
			if from != i {
				right[c][i] = getRight((i+1)%n, col, from) + 1
			}else{
				right[c][i] = n
			}
		}
	}
	return right[c][i]
}

func main() {
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	left = make([][]int, 2)
	right = make([][]int, 2)
	left[0] = make([]int, n)
	right[0] = make([]int, n)
	left[1] = make([]int, n)
	right[1] = make([]int, n)
	for i = 0; i < n; i++ {
		left[0][i] = -1
		left[1][i] = -1
		right[0][i] = -1
		right[1][i] = -1
	}
	for i = 0; i < n; i++ {
		left[0][i] = getLeft(i,'r',(i+1)%n)
		left[1][i] = getLeft(i,'b',(i+1)%n)
		right[0][i] = getRight(i,'r',(i-1+n)%n)
		right[1][i] = getRight(i,'b',(i-1+n)%n)
	}
	//fmt.Println(left[0])
	//fmt.Println(left[1])
	//fmt.Println(right[0])
	//fmt.Println(right[1])
	// ans = max { left[i,col1] + left[i,col2] }
	for i = 0; i < n; i++ {
		if left[0][i] + right[1][(i+1)%n] > ans{
			ans = left[0][i] + right[1][(i+1)%n]
		}
		if left[1][i] +  right[0][(i+1)%n] > ans {
			ans = left[1][i] + right[0][(i+1)%n]
		}
	}
	if ans > n{
		ans = n
	}
	fmt.Println(ans)
}

