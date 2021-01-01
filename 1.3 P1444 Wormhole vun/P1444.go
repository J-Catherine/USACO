package main

import (
	"fmt"
	"sort"
)

var (
	n, ans int
	b      []int
	p      points
)

type point struct {
	x int
	y int
}

type points []point

func (p points) Len() int { return len(p) }
func (p points) Less(i, j int) bool {
	if p[i].y == p[j].y {
		return p[i].x < p[j].x
	}
	return p[i].y < p[j].y
}
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func dfs(step, now, start, flag int, v *[]int, w *[]int) int { //flag==0：以走的方式到达点now   flag==1：从某虫洞到达点now
	//fmt.Println(step,now,start)

	if flag == 0 {
		if (*v)[now] == start {
			return 1
		}
		(*v)[now] = start
		ret := dfs(step+1, b[now], start, 1, v, w)
		if ret == 1 {
			return 1
		}
	} else {
		if (*w)[now] == start {
			return 1
		}
		(*w)[now] = start
		if now+1 <= n && p[now+1].y == p[now].y {
			ret := dfs(step+1, now+1, start, 0, v, w)
			if ret == 1 {
				return 1
			}
		}
	}

	return 0
}

func check() int {
	visited := make([]int, n+1)
	wormhole := make([]int, n+1)
	for i := 1; i <= n; i++ {

		if dfs(0, i, i, 0, &visited, &wormhole) == 1 {
			return 1
		}
	}
	return 0
}

func partner(x int) {
	if x == n+1 {

		//fmt.Println("ans",ans)
		if check() == 1 {
			ans += 1
			if b[2] == 5 {

				//fmt.Println("b",b)
				//for ;true;{;}
			}
		}
		return
	}
	if b[x] == 0 {

		for i := x + 1; i <= n; i++ {
			if b[i] == 0 {
				b[i] = x
				b[x] = i
				partner(x + 1)
				b[i] = 0
				b[x] = 0
			}
		}
	} else {
		partner(x + 1)
	}
}

func main() {
	_, _ = fmt.Scanln(&n)
	p = make([]point, n+1)
	b = make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Scanln(&p[i].x, &p[i].y)
	}
	p[0].x = -1000
	p[0].y = -1000
	sort.Sort(p)
	//fmt.Println(p)
	partner(1)
	fmt.Println(ans)

}
