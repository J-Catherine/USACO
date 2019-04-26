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

func coprime(a, b int) bool {
	for i := 2; i <= 200; i++ {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	return true
}

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
	n := 0
	fmt.Scanln(&n)
	sum := 0
	count := ((n + 1) * n) / 2
	a := make(Numbers, count)

	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if i < j {
				sum += 1
				a[sum].z = float64(i) / float64(j) //分数值，因为i,j是整数，所以要乘1.0000
				a[sum].x = i
				a[sum].y = j
				// fmt.Println(sum,a[sum].z,a[sum].x,a[sum].y)
			}
		}
	}
	sort.Sort(a)
	//fmt.Println(a)
	fmt.Println("0/1")
	for i := range a {
		if a[i].y == 0 {
			continue
		}
		if coprime(a[i].x, a[i].y) {
			fmt.Printf("%d/%d\n", a[i].x, a[i].y)
		}
	}
	fmt.Println("1/1")
}
