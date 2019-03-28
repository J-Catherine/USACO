package main

import (
	"fmt"
	"sort"
)

var(
	n,m,i     int   //需要 n 牛奶，有 m 个提供牛奶的农民
	unitPrice []int // Pi 是农民 i 的牛奶的单价, Ai是农民 i 一天能卖给Marry的牛奶数量
	milk      map[int] int

	val,num,milkNum,pay int
)


func main()  {
	fmt.Scanln(&n,&m)
	milk = make(map[int]int)
	for i=0; i<m; i++{
		fmt.Scanln(&val,&num)
		milk[val] += num
	}
	//fmt.Println(milk)

	unitPrice = make([]int,len(milk)) // 按单价从小到大排
	i=0
	for k := range milk {
		unitPrice[i] = k
		i+= 1
	}
	//fmt.Println(unitPrice)
	sort.Ints(unitPrice)
	//fmt.Println(unitPrice)

	// To perform the opertion you want
	for i=0; milkNum < n; i++ {
		if milkNum + milk[unitPrice[i]] > n {
			pay += (n-milkNum) * unitPrice[i]
			milkNum = n
		} else {
			pay += unitPrice[i]*milk[unitPrice[i]]
			milkNum += milk[unitPrice[i]]
		}
	}
	fmt.Println(pay)
}
