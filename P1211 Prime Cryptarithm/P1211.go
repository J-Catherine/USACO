package main

import (
	"fmt"
)

var (
	n, value1, value2, cnt, flag int
	num, now                            []int
)

//func dfs(x int)  {
//
//}

func judge(value int) {
	for ; value > 0; value /= 10 {
		if now[value%10] != 1 {
			flag = 0
		}
	}
}


func main() {
	fmt.Scanln(&n)
	num = make([]int, n)
	now = make([]int, 10)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
		now[num[i]] = 1
	}
	//fmt.Println(num)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				value1 = num[i]*100 + num[j]*10 + num[k]
				for l := 0; l < n; l++ {
					for m := 0; m < n; m++ {
						value2 = num[l]*10 + num[m]
						if value1*num[m] < 1000 && value1*num[l] < 1000 && value1*value2 < 10000 {
							flag = 1
							judge(value1 * value2)
							judge(value1 * num[m])
							judge(value1 * num[l])
							cnt += flag
						}
					}
				}
			}
		}
	}
	fmt.Println(cnt)

}
