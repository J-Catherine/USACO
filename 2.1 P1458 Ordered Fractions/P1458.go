package main

import (
	"fmt"
	"sort"
)

type Number struct {
	x, y int
	z    float64 //x,y是分子和分母，z是分数值
}
type Numbers []Number

//func gcd(a, b int) int {
//	d := a % b
//	if d == 1 || d == 0{
//		return d
//	} else {
//		return gcd(b,d)
//	}
//}

func (a Numbers) Len() int {
	return len(a)
}
func (a Numbers) Less(i, j int) bool {
	return a[i].z < a[j].z
}
func (a Numbers) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func main() {
	n := 5000
	fmt.Scanln(&n)
	sum := 0
	count := ((n + 1) * n) / 2
	a := make(Numbers, count)

	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			if gcd(i,j) == 1{
				a[sum].z = float64(i) / float64(j)
				a[sum].x = i
				a[sum].y = j
				sum += 1
			}
		}
	}
	a = a[:sum]
	sort.Sort(a)
	//fmt.Println("0/1")
	//for i := range a {
	//	//fmt.Printf("%d/%d\n", a[i].x, a[i].y)
	//}
	//fmt.Println("1/1")
}
